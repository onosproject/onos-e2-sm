<!--
SPDX-FileCopyrightText: 2022-present Intel Corporation
SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
SPDX-License-Identifier: Apache-2.0
-->

## How to deal with decoding problems?
When you implement your SW in ONF’s RIC you may face some issues related to the encoding/decoding of the information 
sent over the wire. Here is a summary of potential problems you may face.

### Incompatibility of APER encodings is the most common issue.
APER encoding is specified in [ITU-T X.691](https://www.itu.int/ITU-T/studygroups/com17/languages/X.691-0207.pdf) and may seem a bit vague.
Reference asn1c tool recommended by O-RAN is Nokia’s distribution of [asn1c](https://github.com/nokia/asn1c).

[Go APER library](https://github.com/onosproject/onos-lib-go/tree/master/pkg/asn1/aper) is a [fork](https://github.com/free5gc/aper) 
of the free5gc project. It was significantly reworked and enhanced. It is fully compatible with Nokia’s asn1c tool. 
It was proved with unit tests for E2AP and E2SMs (KPMv2, RC-PRE, MHO, Test-SM).

#### What may case that?
1. Make sure you've inserted all tags correspondingly to the ASN1 definition of your SM.
   * If you're programming in C, take a look at the source code - you may miss to insert some APER tags.
   * If you’re programming in Golang, make sure that all tags were correctly injected into the resulting Protobuf (in Golang). 
   If you don’t see them (tags), please revisit [tutorial on how to create your own SM](sm-howto.md).

2. Enable DEBUG mode in [Go APER library](https://github.com/onosproject/onos-lib-go/tree/master/pkg/asn1/aper) and do a 
bit by bit analysis (by hand, see next section). It helps to determine what goes wrong in the encoding/decoding process, and helps to understand
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
        int32 choice_a = 1 [json_name = "choiceA"];
        // @inject_tag: aper:"choiceIdx:2"
        int32 choice_b = 2 [json_name = "choiceB"];
    }
};
```

> In E2AP (v1.0 and 2.0) `CHOICE`s are encoded with regard to the canonical ordering. That means that index is not the ordering number, but it is the other information from the ASN.1 \
> definition (i.e., IDs which have `unique` tag). CHOICE index encoding is thus conditional and depends on the value you've put in a `unique` ID. 
> In this case a tag `canonicalOrder` is inserted on top of the CHOICE structure (see [E2AP protobuf](https://github.com/onosproject/onos-e2t/blob/7f0b65294ecd539e15715514ba7a201b7098868f/api/e2ap/v2/e2ap_pdu_contents.proto#L36-L48)) 
> For that purpose a CHOICE map is being created with [protoc-gen-choice](../protoc-gen-choice/README.md) plugin.

`CHOICE` structure can also be extensible and can contain items in its extension. A good example on uch structure could be found [here](https://github.com/onosproject/onos-e2t/blob/7f0b65294ecd539e15715514ba7a201b7098868f/api/e2ap/v2/e2ap_ies.proto#L307-L319). 

* `SEQUENCE` encoding



### Do a bit by bit analysis
Here are some examples


# Should I give a description to each error message which can occur while using Go APER library?
Firstly, read the error message.


## Some useful resources:

[1] [ITU-T X.691 specification](https://www.itu.int/ITU-T/studygroups/com17/languages/X.691-0207.pdf)

[2] [Olivier DuBuisson, ASN.1 Communication between heterogeneous systems (in particular, chapters 15.6 and 20.6.21)](https://www.oss.com/asn1/resources/books-whitepapers-pubs/dubuisson-asn1-book.PDF)

[3] [OSS Novakia ASN.1 notes](https://www.oss.com/asn1/knowledge-center/asn1-java/asn1java-der-support.html)

[4] [A Warm Welcome to ASN.1 and DER (provides a good overview of ASN1 encoding, but not related to APER – may be excluded)](https://letsencrypt.org/docs/a-warm-welcome-to-asn1-and-der/)

[5] [John Larmount, ASN.1 complete](https://www.oss.com/asn1/resources/books-whitepapers-pubs/larmouth-asn1-book.pdf)
