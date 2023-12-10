export ANDROID_HOME="C:\Users\lym\AppData\Local\Android\Sdk"
export ANDROID_NDK_HOME=$ANDROID_HOME/ndk/25.1.8937393

export GOARCH=arm
export GOOS=android
export CGO_ENABLED=1
export CC=$ANDROID_NDK_HOME/toolchains/llvm/prebuilt/windows-x86_64/bin/armv7a-linux-androideabi21-clang
go build -buildmode=c-shared -ldflags="-w -s" -o ../output/android/armv7a/libv2ray.so ../main_lib.go

echo "Build armeabi-v7a success"
