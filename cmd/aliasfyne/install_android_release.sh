#!/bin/bash
#
# Copyright 2021 Aletheia Ware LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -e
set -x

(cd $GOPATH/src/aletheiaware.com/aliasfynego/ui/data/ && ./gen.sh)
go fmt $GOPATH/src/aletheiaware.com/{aliasfynego,aliasfynego/...}
go vet $GOPATH/src/aletheiaware.com/{aliasfynego,aliasfynego/...}
go test $GOPATH/src/aletheiaware.com/{aliasfynego,aliasfynego/...}
(cd $GOPATH/src/aletheiaware.com/aliasfynego/cmd/aliasfyne && fyne release -os android -appID com.aletheiaware.alias -appBuild 1 -appVersion 1.2.3 -keyStore=$GOPATH/src/aletheiaware.com/aliasfynego/private/Alias.keystore -name Alias)
(cd $GOPATH/src/aletheiaware.com/aliasfynego/cmd/aliasfyne && adb install -r -g Alias.apk)
(cd $GOPATH/src/aletheiaware.com/aliasfynego/cmd/aliasfyne && adb logcat -c && adb logcat | tee android.log)
