platforms="windows/386 windows/amd64 windows/arm64 darwin/amd64 darwin/arm64 linux/amd64 linux/arm64"
package_name="go-reverse-proxy"

rm -rf build
mkdir -p build
cd build

for platform in $platforms; do
  echo "Building $platform..."
  platform_split=(${platform//\// })
	GOOS=${platform_split[0]}
	GOARCH=${platform_split[1]}

  output_name=$package_name'-'$GOOS'-'$GOARCH.tar.gz
  binary_name=$package_name

	if [ $GOOS = "windows" ]; then
		binary_name+='.exe'
	fi

  env GOOS=$GOOS GOARCH=$GOARCH go build -o $binary_name ..
  tar -czf $output_name $binary_name
  rm $binary_name
done
