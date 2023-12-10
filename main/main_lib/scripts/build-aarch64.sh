export ANDROID_HOME="C:\Users\lym\AppData\Local\Android\Sdk"
export ANDROID_NDK_HOME=$ANDROID_HOME/ndk/25.1.8937393

export GOARCH=arm64
export GOOS=android
export CGO_ENABLED=1
export CC=$ANDROID_NDK_HOME/toolchains/llvm/prebuilt/windows-x86_64/bin/aarch64-linux-android21-clang
go build -buildmode=c-shared -ldflags="-w -s" -o ../output/android/aarch64/libv2ray.so ../main_lib.go

echo "Build arm64-v8a success"
