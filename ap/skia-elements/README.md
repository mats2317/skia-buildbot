This is an experimental repo of v1 custom elements with a webpack build system.

There are no tests.

Importing
=========

Code from this library should imported under the 'skia-elements' name.
See `pages/index.js` as an example.

If loaded via npm/yarn then importing will just work, i.e.:

    import 'skia-elements/checkbox-sk'

If not loaded via npm/yarn then the `ap` directory will need to be
added in your `webpack.config.js` under [resolve.modules](https://webpack.js.org/configuration/resolve/#resolve-modules),
i.e.:

    module.exports = {
      ...,
      resolve: {
        modules: [path.resolve(__dirname, '..', '..', 'ap'), 'node_modules'],
      },


Disclaimer
==========

This is not an officially supported Google product.
