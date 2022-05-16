<!--
SPDX-FileCopyrightText: 2022-present Intel Corporation
SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
SPDX-License-Identifier: Apache-2.0
-->

# How to deal with encoding/decoding problems?
When you integrate your SW with ONF’s RIC you may face some issues related to the encoding/decoding of the information 
sent over the wire. Here is a summary of potential problems you may face.

### Incompatibility of APER encodings is the most common issue.
APER encoding is specified in [ITU-T X.691](https://www.itu.int/ITU-T/studygroups/com17/languages/X.691-0207.pdf) and may seem to be a bit vague.
In O-RAN, encoding/decoding schema is generated with asn1c tool, which generates a C code. Reference asn1c tool recommended by O-RAN is 
Nokia’s distribution of [asn1c](https://github.com/nokia/asn1c).

[Go APER library](https://github.com/onosproject/onos-lib-go/tree/master/pkg/asn1/aper) is a [fork](https://github.com/free5gc/aper) 
of the free5gc project, which was significantly reworked and enhanced. It is fully compatible with Nokia’s asn1c tool (and thus 
**compliant with O-RAN**), which was proved with unit tests for E2AP and E2SMs (KPMv2, RC-PRE, MHO, Test-SM).

#### What may case incompatibility in APER bytes?
1. Make sure you've inserted all tags in Protobuf, so they correspond to the ASN.1 definition of your SM.
   * If you're programming in C, take a look at the source code - you may miss to insert some APER tags in your code.
   * If you’re programming in Golang, make sure that all tags were correctly injected into the Protobuf (wrapped with Golang, file extension is `.pb.go`). 
   If you don’t see them (tags), please revisit [tutorial on how to create your own SM](sm-howto.md).

2. Enable DEBUG mode in [Go APER library](https://github.com/onosproject/onos-lib-go/tree/master/pkg/asn1/aper) and do a 
bit by bit analysis by hand (see next section). It helps to determine what goes wrong in the encoding/decoding process, and helps to understand
APER encoding/decoding flow better.

### What is the deal with APER tags?
APER tags play a huge role in Go APER library. It is a key for the library to understand how to encode or decode certain structure to or from APER bytes.
A brief summary on APER tags you may find in the [README](https://github.com/onosproject/onos-lib-go/blob/master/pkg/asn1/aper/README.md) of the Go APER library.
We will provide here some basic examples on APER tags usage.

> Please note that APER tags are ***partially*** generated with ONF's distribution of [asn1c](https://github.com/onosproject/asn1c) tool. 
> You still have to insert the majority of the tags manually. This is something to automate in the future.

* `CHOICE` encoding requires inserting indexes. It helps encoder to understand which option has to be encoded or was encoded (in case of the decoding).
```protobuf
message Choice {
    // choice from tes_sm.asn1:65
    oneof choice {
        // @inject_tag: aper:"choiceIdx:1"
        int32 choice_a = 1;
        // @inject_tag: aper:"choiceIdx:2"
        int32 choice_b = 2;
    }
};
```

> In E2AP (v1.0 and 2.0) `CHOICE`s are encoded with regard to the canonical ordering. That means that index is not the ordering number, but it is the other information from the ASN.1 \
> definition (i.e., IDs which have `unique` tag). CHOICE index encoding is thus conditional and depends on the value you've put in a `unique` ID. 
> In this case a tag `canonicalOrder` is inserted on top of the CHOICE structure (see [E2AP protobuf](https://github.com/onosproject/onos-e2t/blob/7f0b65294ecd539e15715514ba7a201b7098868f/api/e2ap/v2/e2ap_pdu_contents.proto#L36-L48)). 
> For that purpose a CHOICE map is being created with [protoc-gen-choice](../protoc-gen-choice/README.md) plugin.

`CHOICE` structure can also be extensible and can contain items in its extension. This requires to put such tags as `choiceExt` and `fromChoiceExt` to indicate that the `CHOICE` can be extended and indicated which items 
belong to its extension A good example of such structure could be found [here](https://github.com/onosproject/onos-e2t/blob/7f0b65294ecd539e15715514ba7a201b7098868f/api/e2ap/v2/e2ap_ies.proto#L307-L319).

* `SEQUENCE` encoding usually requires only single tag to be inserted in case it is needed. It is `valueExt`, which indicates that the `SEQUENCE` can be
extended. If items in the extension are defined, then tag `fromValueExt` is used (see below).
```protobuf
message TopLevelPdu {
   // @inject_tag: aper:"valueExt"
   SequenceExtended value = 1;
}

message SequenceExtended {
   bool se1 = 1;
   // @inject_tag: aper:"optional"
   optional bytes se2 = 2;
   // @inject_tag: aper:"fromValueExt"
   int32 se3 = 3;
   // @inject_tag: aper:"fromValueExt,optional,sizeLB:2,sizeUB:6,sizeExt"
   optional string se4 = 4;
}
```

* `INTEGER` structures in ASN.1 can have constraints. To apply them in encoding, such flags as `valueLB`, `valueUB` and `valueExt` are used.
  * `valueLB:` is used to specify the lowerbound.
  * `valueUB:` is used to specify the upperbound.
  * `valueExt` is used to specify that the `INTEGER` upperbound can be exceeded.
```protobuf
message TestUnconstrainedInt {
   int32 attr_uci_a = 1 [json_name = "attrUciA"];
   int32 attr_uci_b = 2 [json_name = "attrUciB"];
};

message TestConstrainedInt {
   // @inject_tag: aper:"valueLB:10,valueUB:100"
   int32 attr_ci_a = 1 [json_name = "attrCiA"];
   // @inject_tag: aper:"valueLB:255,valueUB:65535"
   int32 attr_ci_b = 2 [json_name = "attrCiB"];
   // @inject_tag: aper:"valueLB:10,valueUB:4294967295"
   int32 attr_ci_c = 3 [json_name = "attrCiC"];
   // @inject_tag: aper:"valueUB:100"
   int32 attr_ci_d = 4 [json_name = "attrCiD"];
   // @inject_tag: aper:"valueLB:10,valueUB:20"
   int32 attr_ci_e = 5 [json_name = "attrCiE"];
   // @inject_tag: aper:"valueLB:10,valueUB:10"
   int32 attr_ci_f = 6 [json_name = "attrCiF"];
   // @inject_tag: aper:"valueLB:10,valueUB:10,valueExt"
   int32 attr_ci_g = 7 [json_name = "attrCiG"];
};
```

* `ENUMERATOR` structures in ASN.1 are treated as a constrained `INTEGER` and thus require to specify lowerbound, upperbound and the extensibility (see above).
```protobuf
message TopLevelPdu {
    // @inject_tag: aper:"valueLB:0,valueUB:5,valueExt"
    TestEnumeratedExtensible value = 1;
};

enum TestEnumeratedExtensible {
   TEST_ENUMERATED_EXTENSIBLE_ENUM1 = 0;
   TEST_ENUMERATED_EXTENSIBLE_ENUM2 = 1;
   TEST_ENUMERATED_EXTENSIBLE_ENUM3 = 2;
   TEST_ENUMERATED_EXTENSIBLE_ENUM4 = 3;
   TEST_ENUMERATED_EXTENSIBLE_ENUM5 = 4;
   TEST_ENUMERATED_EXTENSIBLE_ENUM6 = 5;
};
```

* `PrintableString`, `OCTET STRING` and `BIT STRING` similarly to `INTEGER` can require specifying constraints. Following tags are used:
  * `sizeLB:` is used to specify the lowerbound.
  * `sizeUB:` is used to specify the upperbound.
  * `sizeExt` is used to specify that the `PrintableString`, `OCTET STRING` or `BIT STRING` upperbound can be exceeded.
```protobuf
message BitString {
   asn1.v1.BitString attr_bs1 = 1;
   // @inject_tag: aper:"sizeLB:20,sizeUB:20"
   asn1.v1.BitString attr_bs2 = 2;
   // @inject_tag: aper:"sizeLB:20,sizeUB:20,sizeExt"
   asn1.v1.BitString attr_bs3 = 3;
   // @inject_tag: aper:"sizeLB:0,sizeUB:18"
   asn1.v1.BitString attr_bs4 = 4;
   // @inject_tag: aper:"sizeLB:22,sizeUB:32"
   asn1.v1.BitString attr_bs5 = 5;
   // @inject_tag: aper:"sizeLB:28,sizeUB:32,sizeExt"
   asn1.v1.BitString attr_bs6 = 6;
};

message PrintableString {
   string attr_ps1 = 1;
   // @inject_tag: aper:"sizeLB:2,sizeUB:2"
   string attr_ps2 = 2;
   // @inject_tag: aper:"sizeLB:2,sizeUB:2,sizeExt"
   string attr_ps3 = 3;
   // @inject_tag: aper:"sizeLB:0,sizeUB:3"
   string attr_ps4 = 4;
   // @inject_tag: aper:"sizeLB:2,sizeUB:3"
   string attr_ps5 = 5;
   // @inject_tag: aper:"sizeLB:1,sizeUB:3,sizeExt"
   string attr_ps6 = 6;
};

message OctetString {
   bytes v1 = 1;
   // @inject_tag: aper:"sizeLB:2,sizeUB:2,sizeExt"
   bytes v2 = 3;
   // @inject_tag: aper:"sizeLB:2,sizeUB:5"
   bytes v3 = 3;
}
```

* `SEQUENCE OF` is used to describe lists within ASN.1. It is interpreted in Protobuf as a message with `repeated` label. Lists in ASN.1 can 
also have constraints on its size. Similar tags as for `PrintableString`, `OCTET STRING` and `BIT STRING` are used.
```protobuf
message List {
    // @inject_tag: aper:"sizeLB:0,sizeUB:12"
    repeated float value = 1 [json_name = "value"];
};
```

* ASN.1 definition may have items marked as `OPTIONAL`, which means that the presence of this item is not necessary. Such behavior is treated in a 
Protobuf with `optional` label and `optional` tag.
```protobuf
message Item {
   // @inject_tag: aper:"optional"
   optional int32 item1 = 1 [json_name = "item1"];
};
```

Once all tags are in correct place, produced with Go APER library bytes are guaranteed to be the same as bytes produced with asn1c tool.


### Do a bit by bit analysis
Sometimes, doing a bit by bit analysis may help a lot. General rule is to go through the APER bytes and decode them by hand, i.e., 
verify that all bits correspond to ASN.1 definition. Here are few examples on how to do such analysis:

* An analysis of GlobalKPMnode-ID structure per KPM v2.0.3 specification can be found [here](globalKPMnodeID_analysis.txt)
  * It describes encoding of `CHOICE` structures.
* An analysis of MeasurementLabel structure per KPM v2.0.3 specification can be found [here](measLabel_analysis.txt)
  * It describes encoding of `SEQUENCE` structure with multiple `OPTIONAL` items.
* An analysis of MeasurementData (a part of IndicationMessage) per KPM v2.0.3 specification can be found [here](kpmv2_indication_message_analysis.txt)
  * It describes encoding of lists, i.e., `SEQUENCE OF` structures.
* An analysis of a header of RANfunction-Description per KPMv2.0.3 specification can be found [here](kpmv2_ranfunction_description_analysis.txt)
  * It describes encoding of `SEQUENCE` structure headers.

Please consider tacking a look at the links in "Some useful resources" section. It may help to understand the APER encoding flow.


## General steps on tackling encoding/decoding problems
1. Read error message carefully. You may exceed the boundary or forgot to include an item.
2. Revisit APER tags and make sure they're present in the `.pb.go` file.
3. Do a bit by bit analysis to understand where is the root cause of the problem.


## Some useful resources:

[1] [ITU-T X.691 specification](https://www.itu.int/ITU-T/studygroups/com17/languages/X.691-0207.pdf)

[2] [Olivier DuBuisson, ASN.1 Communication between heterogeneous systems (in particular, chapters 15.6 and 20.6.21)](https://www.oss.com/asn1/resources/books-whitepapers-pubs/dubuisson-asn1-book.PDF)

[3] [OSS Novakia ASN.1 notes](https://www.oss.com/asn1/knowledge-center/asn1-java/asn1java-der-support.html)

[4] [A Warm Welcome to ASN.1 and DER (provides a good overview of ASN1 encoding, but not related to APER – may be excluded)](https://letsencrypt.org/docs/a-warm-welcome-to-asn1-and-der/)

[5] [John Larmount, ASN.1 complete](https://www.oss.com/asn1/resources/books-whitepapers-pubs/larmouth-asn1-book.pdf)
