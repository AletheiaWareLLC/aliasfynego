#!/bin/bash
#
# Copyright 2020 Aletheia Ware LLC
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
ANDROID_NDK_HOME=${ANDROID_HOME}/ndk-bundle/
(cd $GOPATH/src/aletheiaware.com/aliasfynego/cmd && fyne package -os android -appID com.aletheiaware.alias -icon $GOPATH/src/aletheiaware.com/aliasfynego/ui/data/logo.png -name Alias_unaligned)
(cd $GOPATH/src/aletheiaware.com/aliasfynego/cmd && ${ANDROID_HOME}/build-tools/28.0.3/zipalign -f 4 Alias_unaligned.apk Alias.apk)
(cd $GOPATH/src/aletheiaware.com/aliasfynego/cmd && adb install -r -g Alias.apk)
(cd $GOPATH/src/aletheiaware.com/aliasfynego/cmd && adb logcat -c && adb logcat | tee android.log)
