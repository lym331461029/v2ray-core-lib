export ANDROID_HOME="C:\Users\lym\AppData\Local\Android\Sdk"
export ANDROID_NDK_HOME=$ANDROID_HOME/ndk/25.1.8937393

export GOARCH=386
export GOOS=android
export CGO_ENABLED=1
export CC=$ANDROID_NDK_HOME/toolchains/llvm/prebuilt/windows-x86_64/bin/i686-linux-android21-clang
go build -buildmode=c-shared -ldflags="-w -s" -o ../output/android/i686/libv2ray.so ../main_lib.go

echo "Build android i686 success"


export GOARCH=amd64
export GOOS=android
export CGO_ENABLED=1
export CC=$ANDROID_NDK_HOME/toolchains/llvm/prebuilt/windows-x86_64/bin/x86_64-linux-android21-clang
go build -buildmode=c-shared -ldflags="-w -s" -o ../output/android/x86_64/libv2ray.so ../main_lib.go

echo "Build android x86_64 success"
