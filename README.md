# Coralogix AWS Lambda extension

[![goreportcard](https://goreportcard.com/badge/github.com/coralogix/aws-lambda-extension)](https://goreportcard.com/report/github.com/coralogix/aws-lambda-extension)
[![godoc](https://img.shields.io/badge/godoc-reference-brightgreen.svg?style=flat)](https://godoc.org/github.com/coralogix/aws-lambda-extension)
[![license](https://img.shields.io/github/license/coralogix/aws-lambda-extension.svg)](https://raw.githubusercontent.com/coralogix/aws-lambda-extension/master/LICENSE)
[![GitHub issues](https://img.shields.io/github/issues/coralogix/aws-lambda-extension.svg)](https://github.com/coralogix/aws-lambda-extension/issues)
[![GitHub pull requests](https://img.shields.io/github/issues-pr/coralogix/aws-lambda-extension.svg)](https://github.com/coralogix/aws-lambda-extension/pulls)
[![GitHub contributors](https://img.shields.io/github/contributors/coralogix/aws-lambda-extension.svg)](https://github.com/coralogix/aws-lambda-extension/graphs/contributors)

[Coralogix](https://coralogix.com/) is a machine data analytics SaaS platform that drastically improves the delivery & maintenance process for software providers. Using proprietary machine learning algorithms, Coralogix helps over 100 businesses reduce their issue resolution time, improve customer satisfaction and decrease maintenance costs.
The extension provides full integration of lambda function with Coralogix service.

## Installation

1. Open `Serverless Application Repository` service in your AWS Console and find `Coralogix-Lambda-Extension` application.
2. Configure `CompatibleRuntimes` parameter and select [runtimes](https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html) for your purposes (*you can select up to 5 runtimes*).
3. Click to `Deploy` button and wait until extension layer will be created in your account.

## Usage

Add extension layer `coralogix-extension` to your function and define following environment variables:

* **CORALOGIX_PRIVATE_KEY** - A unique ID which represents your company, this Id will be sent to your mail once you register to Coralogix.
* **CORALOGIX_APP_NAME** - Used to separate your environment, e.g. *SuperApp-test/SuperApp-prod*.
* **CORALOGIX_SUB_SYSTEM** - Your application probably has multiple subsystems, for example, *Backend servers, Middleware, Frontend servers etc*.