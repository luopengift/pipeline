#!/usr/bin/env python

import requests
data = {
    "name":"pipeline_test",
    "labels": {
        "usage": "jenkins",
    },
    "stages": [
        {
            "cmd": "cmding"
        },
    ]
}
print requests.post("http://127.0.0.1:18081/pipelines", json=data).json()
