# go-xlive

## Introduction

go-xlive is a tool for working with multitrack audio sessions recorded with a Behringer X-Live addon card

## Why?

Behringer provides a Python 2 sample of how to read a session and convert it to/from their wonky
multi-channel implementation that is limited to 11 minutes per WAV file at 32 channels.

rant: They could have easily avoided this with ExFAT support and 64 bit BWF audio file support but you know..
Behringer.


## Features

### Current
  - Read the SE_LOG.BIN session file and get information about the session

### Planned 
  - Extract the channels from the multichannel file(s) into individual single track files

### Maybe planned
  - Create a new session with their wonky session format and multichannel files

## Installation

coming soon...

## Usage

coming soon...

## License

go-xlive is licensed under the Apache-2.0 license

Copyright (c) 2024 Steve Cross <flip@foxhollow.cc>

>  Licensed under the Apache License, Version 2.0 (the "License");
>  you may not use this file except in compliance with the License.
>  You may obtain a copy of the License at
>
>       http://www.apache.org/licenses/LICENSE-2.0
>
>  Unless required by applicable law or agreed to in writing, software
>  distributed under the License is distributed on an "AS IS" BASIS,
>  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
>  See the License for the specific language governing permissions and
>  limitations under the License.

**NOTE:** This package is being created by reverse engineering the Python 2 samples provided by 
Behringer on the X-Live support page on their site. There is no license attached to that 
download, and no code was directly copied, instead it is only being used as a reference for 
the format of sessionscreatde by the X-Live card. Because of this, I do not believe there to 
be any licensing issues with my releasing my package as Apache-2.0. If I am wrong, I apologize. 
I am not a lawyer and my intention is not to step on anyones toes, but simply provide a useful 
tool for working with the data produced by this device.

## Links

- [Project on GitHub](https://github.com/hairlesshobo/go-xlive/)
- [Project Homepage](https://www.foxhollow.cc/projects/go-xlive/)