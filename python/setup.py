#!/usr/bin/env python3

from setuptools import find_packages, setup

setup(
    name="onos_e2_sm",
    packages=find_packages(),
    python_requires=">=3.7",
    install_requires=["betterproto>=2.0.0b4,<3.0"]
)
