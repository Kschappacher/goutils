module.exports = {
  mode: "production",
  entry: "./src/index.ts",
  devtool: "inline-source-map",
  output: {
    libraryTarget: "module",
  },
  module: {
    rules: [
      {
        test: /\.ts$/,
        include: /src/,
        exclude: /node_modules/,
        loader: "ts-loader",
      },
    ],
  },
  resolve: {
    extensions: [".ts", ".js"],
  },
  experiments: {
    outputModule: true,
  },
};
