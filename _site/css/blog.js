/******/ (function(modules) { // webpackBootstrap
/******/ 	// The module cache
/******/ 	var installedModules = {};
/******/
/******/ 	// The require function
/******/ 	function __webpack_require__(moduleId) {
/******/
/******/ 		// Check if module is in cache
/******/ 		if(installedModules[moduleId]) {
/******/ 			return installedModules[moduleId].exports;
/******/ 		}
/******/ 		// Create a new module (and put it into the cache)
/******/ 		var module = installedModules[moduleId] = {
/******/ 			i: moduleId,
/******/ 			l: false,
/******/ 			exports: {}
/******/ 		};
/******/
/******/ 		// Execute the module function
/******/ 		modules[moduleId].call(module.exports, module, module.exports, __webpack_require__);
/******/
/******/ 		// Flag the module as loaded
/******/ 		module.l = true;
/******/
/******/ 		// Return the exports of the module
/******/ 		return module.exports;
/******/ 	}
/******/
/******/
/******/ 	// expose the modules object (__webpack_modules__)
/******/ 	__webpack_require__.m = modules;
/******/
/******/ 	// expose the module cache
/******/ 	__webpack_require__.c = installedModules;
/******/
/******/ 	// define getter function for harmony exports
/******/ 	__webpack_require__.d = function(exports, name, getter) {
/******/ 		if(!__webpack_require__.o(exports, name)) {
/******/ 			Object.defineProperty(exports, name, {
/******/ 				configurable: false,
/******/ 				enumerable: true,
/******/ 				get: getter
/******/ 			});
/******/ 		}
/******/ 	};
/******/
/******/ 	// getDefaultExport function for compatibility with non-harmony modules
/******/ 	__webpack_require__.n = function(module) {
/******/ 		var getter = module && module.__esModule ?
/******/ 			function getDefault() { return module['default']; } :
/******/ 			function getModuleExports() { return module; };
/******/ 		__webpack_require__.d(getter, 'a', getter);
/******/ 		return getter;
/******/ 	};
/******/
/******/ 	// Object.prototype.hasOwnProperty.call
/******/ 	__webpack_require__.o = function(object, property) { return Object.prototype.hasOwnProperty.call(object, property); };
/******/
/******/ 	// __webpack_public_path__
/******/ 	__webpack_require__.p = "";
/******/
/******/ 	// Load entry module and return exports
/******/ 	return __webpack_require__(__webpack_require__.s = 5);
/******/ })
/************************************************************************/
/******/ ([
/* 0 */
/***/ (function(module, exports) {

/*
	MIT License http://www.opensource.org/licenses/mit-license.php
	Author Tobias Koppers @sokra
*/
// css base code, injected by the css-loader
module.exports = function(useSourceMap) {
	var list = [];

	// return the list of modules as css string
	list.toString = function toString() {
		return this.map(function (item) {
			var content = cssWithMappingToString(item, useSourceMap);
			if(item[2]) {
				return "@media " + item[2] + "{" + content + "}";
			} else {
				return content;
			}
		}).join("");
	};

	// import a list of modules into the list
	list.i = function(modules, mediaQuery) {
		if(typeof modules === "string")
			modules = [[null, modules, ""]];
		var alreadyImportedModules = {};
		for(var i = 0; i < this.length; i++) {
			var id = this[i][0];
			if(typeof id === "number")
				alreadyImportedModules[id] = true;
		}
		for(i = 0; i < modules.length; i++) {
			var item = modules[i];
			// skip already imported module
			// this implementation is not 100% perfect for weird media query combinations
			//  when a module is imported multiple times with different media queries.
			//  I hope this will never occur (Hey this way we have smaller bundles)
			if(typeof item[0] !== "number" || !alreadyImportedModules[item[0]]) {
				if(mediaQuery && !item[2]) {
					item[2] = mediaQuery;
				} else if(mediaQuery) {
					item[2] = "(" + item[2] + ") and (" + mediaQuery + ")";
				}
				list.push(item);
			}
		}
	};
	return list;
};

function cssWithMappingToString(item, useSourceMap) {
	var content = item[1] || '';
	var cssMapping = item[3];
	if (!cssMapping) {
		return content;
	}

	if (useSourceMap && typeof btoa === 'function') {
		var sourceMapping = toComment(cssMapping);
		var sourceURLs = cssMapping.sources.map(function (source) {
			return '/*# sourceURL=' + cssMapping.sourceRoot + source + ' */'
		});

		return [content].concat(sourceURLs).concat([sourceMapping]).join('\n');
	}

	return [content].join('\n');
}

// Adapted from convert-source-map (MIT)
function toComment(sourceMap) {
	// eslint-disable-next-line no-undef
	var base64 = btoa(unescape(encodeURIComponent(JSON.stringify(sourceMap))));
	var data = 'sourceMappingURL=data:application/json;charset=utf-8;base64,' + base64;

	return '/*# ' + data + ' */';
}


/***/ }),
/* 1 */,
/* 2 */,
/* 3 */,
/* 4 */,
/* 5 */
/***/ (function(module, exports, __webpack_require__) {

module.exports = __webpack_require__(6);


/***/ }),
/* 6 */
/***/ (function(module, exports, __webpack_require__) {

// style-loader: Adds some css to the DOM by adding a <style> tag

// load the styles
var content = __webpack_require__(7);
if(typeof content === 'string') content = [[module.i, content, '']];
// Prepare cssTransformation
var transform;

var options = {}
options.transform = transform
// add the styles to the DOM
var update = __webpack_require__(67)(content, options);
if(content.locals) module.exports = content.locals;
// Hot Module Replacement
if(false) {
	// When the styles change, update the <style> tags
	if(!content.locals) {
		module.hot.accept("!!../../node_modules/css-loader/index.js!../../node_modules/sass-loader/lib/loader.js!./blog.scss", function() {
			var newContent = require("!!../../node_modules/css-loader/index.js!../../node_modules/sass-loader/lib/loader.js!./blog.scss");
			if(typeof newContent === 'string') newContent = [[module.id, newContent, '']];
			update(newContent);
		});
	}
	// When the module is disposed, remove the <style> tags
	module.hot.dispose(function() { update(); });
}

/***/ }),
/* 7 */
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__(0)(false);
// imports
exports.i(__webpack_require__(8), "");

// module
exports.push([module.i, "\n", ""]);

// exports


/***/ }),
/* 8 */
/***/ (function(module, exports, __webpack_require__) {

var escape = __webpack_require__(9);
exports = module.exports = __webpack_require__(0)(false);
// imports


// module
exports.push([module.i, "@font-face {\n  font-family: 'KaTeX_AMS';\n  src: url(" + escape(__webpack_require__(10)) + ") format('woff2'), url(" + escape(__webpack_require__(11)) + ") format('woff'), url(" + escape(__webpack_require__(12)) + ") format('truetype');\n  font-weight: normal;\n  font-style: normal;\n}\n@font-face {\n  font-family: 'KaTeX_Caligraphic';\n  src: url(" + escape(__webpack_require__(13)) + ") format('woff2'), url(" + escape(__webpack_require__(14)) + ") format('woff'), url(" + escape(__webpack_require__(15)) + ") format('truetype');\n  font-weight: bold;\n  font-style: normal;\n}\n@font-face {\n  font-family: 'KaTeX_Caligraphic';\n  src: url(" + escape(__webpack_require__(16)) + ") format('woff2'), url(" + escape(__webpack_require__(17)) + ") format('woff'), url(" + escape(__webpack_require__(18)) + ") format('truetype');\n  font-weight: normal;\n  font-style: normal;\n}\n@font-face {\n  font-family: 'KaTeX_Fraktur';\n  src: url(" + escape(__webpack_require__(19)) + ") format('woff2'), url(" + escape(__webpack_require__(20)) + ") format('woff'), url(" + escape(__webpack_require__(21)) + ") format('truetype');\n  font-weight: bold;\n  font-style: normal;\n}\n@font-face {\n  font-family: 'KaTeX_Fraktur';\n  src: url(" + escape(__webpack_require__(22)) + ") format('woff2'), url(" + escape(__webpack_require__(23)) + ") format('woff'), url(" + escape(__webpack_require__(24)) + ") format('truetype');\n  font-weight: normal;\n  font-style: normal;\n}\n@font-face {\n  font-family: 'KaTeX_Main';\n  src: url(" + escape(__webpack_require__(25)) + ") format('woff2'), url(" + escape(__webpack_require__(26)) + ") format('woff'), url(" + escape(__webpack_require__(27)) + ") format('truetype');\n  font-weight: bold;\n  font-style: normal;\n}\n@font-face {\n  font-family: 'KaTeX_Main';\n  src: url(" + escape(__webpack_require__(28)) + ") format('woff2'), url(" + escape(__webpack_require__(29)) + ") format('woff'), url(" + escape(__webpack_require__(30)) + ") format('truetype');\n  font-weight: bold;\n  font-style: italic;\n}\n@font-face {\n  font-family: 'KaTeX_Main';\n  src: url(" + escape(__webpack_require__(31)) + ") format('woff2'), url(" + escape(__webpack_require__(32)) + ") format('woff'), url(" + escape(__webpack_require__(33)) + ") format('truetype');\n  font-weight: normal;\n  font-style: italic;\n}\n@font-face {\n  font-family: 'KaTeX_Main';\n  src: url(" + escape(__webpack_require__(34)) + ") format('woff2'), url(" + escape(__webpack_require__(35)) + ") format('woff'), url(" + escape(__webpack_require__(36)) + ") format('truetype');\n  font-weight: normal;\n  font-style: normal;\n}\n@font-face {\n  font-family: 'KaTeX_Math';\n  src: url(" + escape(__webpack_require__(37)) + ") format('woff2'), url(" + escape(__webpack_require__(38)) + ") format('woff'), url(" + escape(__webpack_require__(39)) + ") format('truetype');\n  font-weight: normal;\n  font-style: italic;\n}\n@font-face {\n  font-family: 'KaTeX_SansSerif';\n  src: url(" + escape(__webpack_require__(40)) + ") format('woff2'), url(" + escape(__webpack_require__(41)) + ") format('woff'), url(" + escape(__webpack_require__(42)) + ") format('truetype');\n  font-weight: bold;\n  font-style: normal;\n}\n@font-face {\n  font-family: 'KaTeX_SansSerif';\n  src: url(" + escape(__webpack_require__(43)) + ") format('woff2'), url(" + escape(__webpack_require__(44)) + ") format('woff'), url(" + escape(__webpack_require__(45)) + ") format('truetype');\n  font-weight: normal;\n  font-style: italic;\n}\n@font-face {\n  font-family: 'KaTeX_SansSerif';\n  src: url(" + escape(__webpack_require__(46)) + ") format('woff2'), url(" + escape(__webpack_require__(47)) + ") format('woff'), url(" + escape(__webpack_require__(48)) + ") format('truetype');\n  font-weight: normal;\n  font-style: normal;\n}\n@font-face {\n  font-family: 'KaTeX_Script';\n  src: url(" + escape(__webpack_require__(49)) + ") format('woff2'), url(" + escape(__webpack_require__(50)) + ") format('woff'), url(" + escape(__webpack_require__(51)) + ") format('truetype');\n  font-weight: normal;\n  font-style: normal;\n}\n@font-face {\n  font-family: 'KaTeX_Size1';\n  src: url(" + escape(__webpack_require__(52)) + ") format('woff2'), url(" + escape(__webpack_require__(53)) + ") format('woff'), url(" + escape(__webpack_require__(54)) + ") format('truetype');\n  font-weight: normal;\n  font-style: normal;\n}\n@font-face {\n  font-family: 'KaTeX_Size2';\n  src: url(" + escape(__webpack_require__(55)) + ") format('woff2'), url(" + escape(__webpack_require__(56)) + ") format('woff'), url(" + escape(__webpack_require__(57)) + ") format('truetype');\n  font-weight: normal;\n  font-style: normal;\n}\n@font-face {\n  font-family: 'KaTeX_Size3';\n  src: url(" + escape(__webpack_require__(58)) + ") format('woff2'), url(" + escape(__webpack_require__(59)) + ") format('woff'), url(" + escape(__webpack_require__(60)) + ") format('truetype');\n  font-weight: normal;\n  font-style: normal;\n}\n@font-face {\n  font-family: 'KaTeX_Size4';\n  src: url(" + escape(__webpack_require__(61)) + ") format('woff2'), url(" + escape(__webpack_require__(62)) + ") format('woff'), url(" + escape(__webpack_require__(63)) + ") format('truetype');\n  font-weight: normal;\n  font-style: normal;\n}\n@font-face {\n  font-family: 'KaTeX_Typewriter';\n  src: url(" + escape(__webpack_require__(64)) + ") format('woff2'), url(" + escape(__webpack_require__(65)) + ") format('woff'), url(" + escape(__webpack_require__(66)) + ") format('truetype');\n  font-weight: normal;\n  font-style: normal;\n}\n.katex-display {\n  display: block;\n  margin: 1em 0;\n  text-align: center;\n}\n.katex-display > .katex {\n  display: inline-block;\n  text-align: initial;\n}\n.katex {\n  font: normal 1.21em KaTeX_Main, Times New Roman, serif;\n  line-height: 1.2;\n  white-space: nowrap;\n  text-indent: 0;\n  text-rendering: auto;\n}\n.katex * {\n  -ms-high-contrast-adjust: none !important;\n}\n.katex .katex-html {\n  display: inline-block;\n}\n.katex .katex-mathml {\n  position: absolute;\n  clip: rect(1px, 1px, 1px, 1px);\n  padding: 0;\n  border: 0;\n  height: 1px;\n  width: 1px;\n  overflow: hidden;\n}\n.katex .base {\n  position: relative;\n  display: inline-block;\n}\n.katex .strut {\n  display: inline-block;\n}\n.katex .textbf {\n  font-weight: bold;\n}\n.katex .textit {\n  font-style: italic;\n}\n.katex .textrm {\n  font-family: KaTeX_Main;\n}\n.katex .textsf {\n  font-family: KaTeX_SansSerif;\n}\n.katex .texttt {\n  font-family: KaTeX_Typewriter;\n}\n.katex .mathit {\n  font-family: KaTeX_Math;\n  font-style: italic;\n}\n.katex .mathrm {\n  font-style: normal;\n}\n.katex .mathbf {\n  font-family: KaTeX_Main;\n  font-weight: bold;\n}\n.katex .boldsymbol {\n  font-family: KaTeX_Math;\n  font-weight: bold;\n  font-style: italic;\n}\n.katex .amsrm {\n  font-family: KaTeX_AMS;\n}\n.katex .mathbb {\n  font-family: KaTeX_AMS;\n}\n.katex .mathcal {\n  font-family: KaTeX_Caligraphic;\n}\n.katex .mathfrak {\n  font-family: KaTeX_Fraktur;\n}\n.katex .mathtt {\n  font-family: KaTeX_Typewriter;\n}\n.katex .mathscr {\n  font-family: KaTeX_Script;\n}\n.katex .mathsf {\n  font-family: KaTeX_SansSerif;\n}\n.katex .mainit {\n  font-family: KaTeX_Main;\n  font-style: italic;\n}\n.katex .mainrm {\n  font-family: KaTeX_Main;\n  font-style: normal;\n}\n.katex .vlist-t {\n  display: inline-table;\n  table-layout: fixed;\n}\n.katex .vlist-r {\n  display: table-row;\n}\n.katex .vlist {\n  display: table-cell;\n  vertical-align: bottom;\n  position: relative;\n}\n.katex .vlist > span {\n  display: block;\n  height: 0;\n  position: relative;\n}\n.katex .vlist > span > span {\n  display: inline-block;\n}\n.katex .vlist > span > .pstrut {\n  overflow: hidden;\n  width: 0;\n}\n.katex .vlist-t2 {\n  margin-right: -2px;\n}\n.katex .vlist-s {\n  display: table-cell;\n  vertical-align: bottom;\n  font-size: 1px;\n  width: 2px;\n}\n.katex .msupsub {\n  text-align: left;\n}\n.katex .mfrac > span > span {\n  text-align: center;\n}\n.katex .mfrac .frac-line {\n  display: inline-block;\n  width: 100%;\n}\n.katex .mspace {\n  display: inline-block;\n}\n.katex .mspace.negativethinspace {\n  margin-left: -0.16667em;\n}\n.katex .mspace.muspace {\n  width: 0.055556em;\n}\n.katex .mspace.thinspace {\n  width: 0.16667em;\n}\n.katex .mspace.negativemediumspace {\n  margin-left: -0.22222em;\n}\n.katex .mspace.mediumspace {\n  width: 0.22222em;\n}\n.katex .mspace.thickspace {\n  width: 0.27778em;\n}\n.katex .mspace.sixmuspace {\n  width: 0.333333em;\n}\n.katex .mspace.eightmuspace {\n  width: 0.444444em;\n}\n.katex .mspace.enspace {\n  width: 0.5em;\n}\n.katex .mspace.twelvemuspace {\n  width: 0.666667em;\n}\n.katex .mspace.quad {\n  width: 1em;\n}\n.katex .mspace.qquad {\n  width: 2em;\n}\n.katex .llap,\n.katex .rlap,\n.katex .clap {\n  width: 0;\n  position: relative;\n}\n.katex .llap > .inner,\n.katex .rlap > .inner,\n.katex .clap > .inner {\n  position: absolute;\n}\n.katex .llap > .fix,\n.katex .rlap > .fix,\n.katex .clap > .fix {\n  display: inline-block;\n}\n.katex .llap > .inner {\n  right: 0;\n}\n.katex .rlap > .inner,\n.katex .clap > .inner {\n  left: 0;\n}\n.katex .clap > .inner > span {\n  margin-left: -50%;\n  margin-right: 50%;\n}\n.katex .rule {\n  display: inline-block;\n  border: solid 0;\n  position: relative;\n}\n.katex .overline .overline-line,\n.katex .underline .underline-line {\n  display: inline-block;\n  width: 100%;\n}\n.katex .sqrt > .root {\n  margin-left: 0.27777778em;\n  margin-right: -0.55555556em;\n}\n.katex .sizing,\n.katex .fontsize-ensurer {\n  display: inline-block;\n}\n.katex .sizing.reset-size1.size1,\n.katex .fontsize-ensurer.reset-size1.size1 {\n  font-size: 1em;\n}\n.katex .sizing.reset-size1.size2,\n.katex .fontsize-ensurer.reset-size1.size2 {\n  font-size: 1.2em;\n}\n.katex .sizing.reset-size1.size3,\n.katex .fontsize-ensurer.reset-size1.size3 {\n  font-size: 1.4em;\n}\n.katex .sizing.reset-size1.size4,\n.katex .fontsize-ensurer.reset-size1.size4 {\n  font-size: 1.6em;\n}\n.katex .sizing.reset-size1.size5,\n.katex .fontsize-ensurer.reset-size1.size5 {\n  font-size: 1.8em;\n}\n.katex .sizing.reset-size1.size6,\n.katex .fontsize-ensurer.reset-size1.size6 {\n  font-size: 2em;\n}\n.katex .sizing.reset-size1.size7,\n.katex .fontsize-ensurer.reset-size1.size7 {\n  font-size: 2.4em;\n}\n.katex .sizing.reset-size1.size8,\n.katex .fontsize-ensurer.reset-size1.size8 {\n  font-size: 2.88em;\n}\n.katex .sizing.reset-size1.size9,\n.katex .fontsize-ensurer.reset-size1.size9 {\n  font-size: 3.456em;\n}\n.katex .sizing.reset-size1.size10,\n.katex .fontsize-ensurer.reset-size1.size10 {\n  font-size: 4.148em;\n}\n.katex .sizing.reset-size1.size11,\n.katex .fontsize-ensurer.reset-size1.size11 {\n  font-size: 4.976em;\n}\n.katex .sizing.reset-size2.size1,\n.katex .fontsize-ensurer.reset-size2.size1 {\n  font-size: 0.83333333em;\n}\n.katex .sizing.reset-size2.size2,\n.katex .fontsize-ensurer.reset-size2.size2 {\n  font-size: 1em;\n}\n.katex .sizing.reset-size2.size3,\n.katex .fontsize-ensurer.reset-size2.size3 {\n  font-size: 1.16666667em;\n}\n.katex .sizing.reset-size2.size4,\n.katex .fontsize-ensurer.reset-size2.size4 {\n  font-size: 1.33333333em;\n}\n.katex .sizing.reset-size2.size5,\n.katex .fontsize-ensurer.reset-size2.size5 {\n  font-size: 1.5em;\n}\n.katex .sizing.reset-size2.size6,\n.katex .fontsize-ensurer.reset-size2.size6 {\n  font-size: 1.66666667em;\n}\n.katex .sizing.reset-size2.size7,\n.katex .fontsize-ensurer.reset-size2.size7 {\n  font-size: 2em;\n}\n.katex .sizing.reset-size2.size8,\n.katex .fontsize-ensurer.reset-size2.size8 {\n  font-size: 2.4em;\n}\n.katex .sizing.reset-size2.size9,\n.katex .fontsize-ensurer.reset-size2.size9 {\n  font-size: 2.88em;\n}\n.katex .sizing.reset-size2.size10,\n.katex .fontsize-ensurer.reset-size2.size10 {\n  font-size: 3.45666667em;\n}\n.katex .sizing.reset-size2.size11,\n.katex .fontsize-ensurer.reset-size2.size11 {\n  font-size: 4.14666667em;\n}\n.katex .sizing.reset-size3.size1,\n.katex .fontsize-ensurer.reset-size3.size1 {\n  font-size: 0.71428571em;\n}\n.katex .sizing.reset-size3.size2,\n.katex .fontsize-ensurer.reset-size3.size2 {\n  font-size: 0.85714286em;\n}\n.katex .sizing.reset-size3.size3,\n.katex .fontsize-ensurer.reset-size3.size3 {\n  font-size: 1em;\n}\n.katex .sizing.reset-size3.size4,\n.katex .fontsize-ensurer.reset-size3.size4 {\n  font-size: 1.14285714em;\n}\n.katex .sizing.reset-size3.size5,\n.katex .fontsize-ensurer.reset-size3.size5 {\n  font-size: 1.28571429em;\n}\n.katex .sizing.reset-size3.size6,\n.katex .fontsize-ensurer.reset-size3.size6 {\n  font-size: 1.42857143em;\n}\n.katex .sizing.reset-size3.size7,\n.katex .fontsize-ensurer.reset-size3.size7 {\n  font-size: 1.71428571em;\n}\n.katex .sizing.reset-size3.size8,\n.katex .fontsize-ensurer.reset-size3.size8 {\n  font-size: 2.05714286em;\n}\n.katex .sizing.reset-size3.size9,\n.katex .fontsize-ensurer.reset-size3.size9 {\n  font-size: 2.46857143em;\n}\n.katex .sizing.reset-size3.size10,\n.katex .fontsize-ensurer.reset-size3.size10 {\n  font-size: 2.96285714em;\n}\n.katex .sizing.reset-size3.size11,\n.katex .fontsize-ensurer.reset-size3.size11 {\n  font-size: 3.55428571em;\n}\n.katex .sizing.reset-size4.size1,\n.katex .fontsize-ensurer.reset-size4.size1 {\n  font-size: 0.625em;\n}\n.katex .sizing.reset-size4.size2,\n.katex .fontsize-ensurer.reset-size4.size2 {\n  font-size: 0.75em;\n}\n.katex .sizing.reset-size4.size3,\n.katex .fontsize-ensurer.reset-size4.size3 {\n  font-size: 0.875em;\n}\n.katex .sizing.reset-size4.size4,\n.katex .fontsize-ensurer.reset-size4.size4 {\n  font-size: 1em;\n}\n.katex .sizing.reset-size4.size5,\n.katex .fontsize-ensurer.reset-size4.size5 {\n  font-size: 1.125em;\n}\n.katex .sizing.reset-size4.size6,\n.katex .fontsize-ensurer.reset-size4.size6 {\n  font-size: 1.25em;\n}\n.katex .sizing.reset-size4.size7,\n.katex .fontsize-ensurer.reset-size4.size7 {\n  font-size: 1.5em;\n}\n.katex .sizing.reset-size4.size8,\n.katex .fontsize-ensurer.reset-size4.size8 {\n  font-size: 1.8em;\n}\n.katex .sizing.reset-size4.size9,\n.katex .fontsize-ensurer.reset-size4.size9 {\n  font-size: 2.16em;\n}\n.katex .sizing.reset-size4.size10,\n.katex .fontsize-ensurer.reset-size4.size10 {\n  font-size: 2.5925em;\n}\n.katex .sizing.reset-size4.size11,\n.katex .fontsize-ensurer.reset-size4.size11 {\n  font-size: 3.11em;\n}\n.katex .sizing.reset-size5.size1,\n.katex .fontsize-ensurer.reset-size5.size1 {\n  font-size: 0.55555556em;\n}\n.katex .sizing.reset-size5.size2,\n.katex .fontsize-ensurer.reset-size5.size2 {\n  font-size: 0.66666667em;\n}\n.katex .sizing.reset-size5.size3,\n.katex .fontsize-ensurer.reset-size5.size3 {\n  font-size: 0.77777778em;\n}\n.katex .sizing.reset-size5.size4,\n.katex .fontsize-ensurer.reset-size5.size4 {\n  font-size: 0.88888889em;\n}\n.katex .sizing.reset-size5.size5,\n.katex .fontsize-ensurer.reset-size5.size5 {\n  font-size: 1em;\n}\n.katex .sizing.reset-size5.size6,\n.katex .fontsize-ensurer.reset-size5.size6 {\n  font-size: 1.11111111em;\n}\n.katex .sizing.reset-size5.size7,\n.katex .fontsize-ensurer.reset-size5.size7 {\n  font-size: 1.33333333em;\n}\n.katex .sizing.reset-size5.size8,\n.katex .fontsize-ensurer.reset-size5.size8 {\n  font-size: 1.6em;\n}\n.katex .sizing.reset-size5.size9,\n.katex .fontsize-ensurer.reset-size5.size9 {\n  font-size: 1.92em;\n}\n.katex .sizing.reset-size5.size10,\n.katex .fontsize-ensurer.reset-size5.size10 {\n  font-size: 2.30444444em;\n}\n.katex .sizing.reset-size5.size11,\n.katex .fontsize-ensurer.reset-size5.size11 {\n  font-size: 2.76444444em;\n}\n.katex .sizing.reset-size6.size1,\n.katex .fontsize-ensurer.reset-size6.size1 {\n  font-size: 0.5em;\n}\n.katex .sizing.reset-size6.size2,\n.katex .fontsize-ensurer.reset-size6.size2 {\n  font-size: 0.6em;\n}\n.katex .sizing.reset-size6.size3,\n.katex .fontsize-ensurer.reset-size6.size3 {\n  font-size: 0.7em;\n}\n.katex .sizing.reset-size6.size4,\n.katex .fontsize-ensurer.reset-size6.size4 {\n  font-size: 0.8em;\n}\n.katex .sizing.reset-size6.size5,\n.katex .fontsize-ensurer.reset-size6.size5 {\n  font-size: 0.9em;\n}\n.katex .sizing.reset-size6.size6,\n.katex .fontsize-ensurer.reset-size6.size6 {\n  font-size: 1em;\n}\n.katex .sizing.reset-size6.size7,\n.katex .fontsize-ensurer.reset-size6.size7 {\n  font-size: 1.2em;\n}\n.katex .sizing.reset-size6.size8,\n.katex .fontsize-ensurer.reset-size6.size8 {\n  font-size: 1.44em;\n}\n.katex .sizing.reset-size6.size9,\n.katex .fontsize-ensurer.reset-size6.size9 {\n  font-size: 1.728em;\n}\n.katex .sizing.reset-size6.size10,\n.katex .fontsize-ensurer.reset-size6.size10 {\n  font-size: 2.074em;\n}\n.katex .sizing.reset-size6.size11,\n.katex .fontsize-ensurer.reset-size6.size11 {\n  font-size: 2.488em;\n}\n.katex .sizing.reset-size7.size1,\n.katex .fontsize-ensurer.reset-size7.size1 {\n  font-size: 0.41666667em;\n}\n.katex .sizing.reset-size7.size2,\n.katex .fontsize-ensurer.reset-size7.size2 {\n  font-size: 0.5em;\n}\n.katex .sizing.reset-size7.size3,\n.katex .fontsize-ensurer.reset-size7.size3 {\n  font-size: 0.58333333em;\n}\n.katex .sizing.reset-size7.size4,\n.katex .fontsize-ensurer.reset-size7.size4 {\n  font-size: 0.66666667em;\n}\n.katex .sizing.reset-size7.size5,\n.katex .fontsize-ensurer.reset-size7.size5 {\n  font-size: 0.75em;\n}\n.katex .sizing.reset-size7.size6,\n.katex .fontsize-ensurer.reset-size7.size6 {\n  font-size: 0.83333333em;\n}\n.katex .sizing.reset-size7.size7,\n.katex .fontsize-ensurer.reset-size7.size7 {\n  font-size: 1em;\n}\n.katex .sizing.reset-size7.size8,\n.katex .fontsize-ensurer.reset-size7.size8 {\n  font-size: 1.2em;\n}\n.katex .sizing.reset-size7.size9,\n.katex .fontsize-ensurer.reset-size7.size9 {\n  font-size: 1.44em;\n}\n.katex .sizing.reset-size7.size10,\n.katex .fontsize-ensurer.reset-size7.size10 {\n  font-size: 1.72833333em;\n}\n.katex .sizing.reset-size7.size11,\n.katex .fontsize-ensurer.reset-size7.size11 {\n  font-size: 2.07333333em;\n}\n.katex .sizing.reset-size8.size1,\n.katex .fontsize-ensurer.reset-size8.size1 {\n  font-size: 0.34722222em;\n}\n.katex .sizing.reset-size8.size2,\n.katex .fontsize-ensurer.reset-size8.size2 {\n  font-size: 0.41666667em;\n}\n.katex .sizing.reset-size8.size3,\n.katex .fontsize-ensurer.reset-size8.size3 {\n  font-size: 0.48611111em;\n}\n.katex .sizing.reset-size8.size4,\n.katex .fontsize-ensurer.reset-size8.size4 {\n  font-size: 0.55555556em;\n}\n.katex .sizing.reset-size8.size5,\n.katex .fontsize-ensurer.reset-size8.size5 {\n  font-size: 0.625em;\n}\n.katex .sizing.reset-size8.size6,\n.katex .fontsize-ensurer.reset-size8.size6 {\n  font-size: 0.69444444em;\n}\n.katex .sizing.reset-size8.size7,\n.katex .fontsize-ensurer.reset-size8.size7 {\n  font-size: 0.83333333em;\n}\n.katex .sizing.reset-size8.size8,\n.katex .fontsize-ensurer.reset-size8.size8 {\n  font-size: 1em;\n}\n.katex .sizing.reset-size8.size9,\n.katex .fontsize-ensurer.reset-size8.size9 {\n  font-size: 1.2em;\n}\n.katex .sizing.reset-size8.size10,\n.katex .fontsize-ensurer.reset-size8.size10 {\n  font-size: 1.44027778em;\n}\n.katex .sizing.reset-size8.size11,\n.katex .fontsize-ensurer.reset-size8.size11 {\n  font-size: 1.72777778em;\n}\n.katex .sizing.reset-size9.size1,\n.katex .fontsize-ensurer.reset-size9.size1 {\n  font-size: 0.28935185em;\n}\n.katex .sizing.reset-size9.size2,\n.katex .fontsize-ensurer.reset-size9.size2 {\n  font-size: 0.34722222em;\n}\n.katex .sizing.reset-size9.size3,\n.katex .fontsize-ensurer.reset-size9.size3 {\n  font-size: 0.40509259em;\n}\n.katex .sizing.reset-size9.size4,\n.katex .fontsize-ensurer.reset-size9.size4 {\n  font-size: 0.46296296em;\n}\n.katex .sizing.reset-size9.size5,\n.katex .fontsize-ensurer.reset-size9.size5 {\n  font-size: 0.52083333em;\n}\n.katex .sizing.reset-size9.size6,\n.katex .fontsize-ensurer.reset-size9.size6 {\n  font-size: 0.5787037em;\n}\n.katex .sizing.reset-size9.size7,\n.katex .fontsize-ensurer.reset-size9.size7 {\n  font-size: 0.69444444em;\n}\n.katex .sizing.reset-size9.size8,\n.katex .fontsize-ensurer.reset-size9.size8 {\n  font-size: 0.83333333em;\n}\n.katex .sizing.reset-size9.size9,\n.katex .fontsize-ensurer.reset-size9.size9 {\n  font-size: 1em;\n}\n.katex .sizing.reset-size9.size10,\n.katex .fontsize-ensurer.reset-size9.size10 {\n  font-size: 1.20023148em;\n}\n.katex .sizing.reset-size9.size11,\n.katex .fontsize-ensurer.reset-size9.size11 {\n  font-size: 1.43981481em;\n}\n.katex .sizing.reset-size10.size1,\n.katex .fontsize-ensurer.reset-size10.size1 {\n  font-size: 0.24108004em;\n}\n.katex .sizing.reset-size10.size2,\n.katex .fontsize-ensurer.reset-size10.size2 {\n  font-size: 0.28929605em;\n}\n.katex .sizing.reset-size10.size3,\n.katex .fontsize-ensurer.reset-size10.size3 {\n  font-size: 0.33751205em;\n}\n.katex .sizing.reset-size10.size4,\n.katex .fontsize-ensurer.reset-size10.size4 {\n  font-size: 0.38572806em;\n}\n.katex .sizing.reset-size10.size5,\n.katex .fontsize-ensurer.reset-size10.size5 {\n  font-size: 0.43394407em;\n}\n.katex .sizing.reset-size10.size6,\n.katex .fontsize-ensurer.reset-size10.size6 {\n  font-size: 0.48216008em;\n}\n.katex .sizing.reset-size10.size7,\n.katex .fontsize-ensurer.reset-size10.size7 {\n  font-size: 0.57859209em;\n}\n.katex .sizing.reset-size10.size8,\n.katex .fontsize-ensurer.reset-size10.size8 {\n  font-size: 0.69431051em;\n}\n.katex .sizing.reset-size10.size9,\n.katex .fontsize-ensurer.reset-size10.size9 {\n  font-size: 0.83317261em;\n}\n.katex .sizing.reset-size10.size10,\n.katex .fontsize-ensurer.reset-size10.size10 {\n  font-size: 1em;\n}\n.katex .sizing.reset-size10.size11,\n.katex .fontsize-ensurer.reset-size10.size11 {\n  font-size: 1.19961427em;\n}\n.katex .sizing.reset-size11.size1,\n.katex .fontsize-ensurer.reset-size11.size1 {\n  font-size: 0.20096463em;\n}\n.katex .sizing.reset-size11.size2,\n.katex .fontsize-ensurer.reset-size11.size2 {\n  font-size: 0.24115756em;\n}\n.katex .sizing.reset-size11.size3,\n.katex .fontsize-ensurer.reset-size11.size3 {\n  font-size: 0.28135048em;\n}\n.katex .sizing.reset-size11.size4,\n.katex .fontsize-ensurer.reset-size11.size4 {\n  font-size: 0.32154341em;\n}\n.katex .sizing.reset-size11.size5,\n.katex .fontsize-ensurer.reset-size11.size5 {\n  font-size: 0.36173633em;\n}\n.katex .sizing.reset-size11.size6,\n.katex .fontsize-ensurer.reset-size11.size6 {\n  font-size: 0.40192926em;\n}\n.katex .sizing.reset-size11.size7,\n.katex .fontsize-ensurer.reset-size11.size7 {\n  font-size: 0.48231511em;\n}\n.katex .sizing.reset-size11.size8,\n.katex .fontsize-ensurer.reset-size11.size8 {\n  font-size: 0.57877814em;\n}\n.katex .sizing.reset-size11.size9,\n.katex .fontsize-ensurer.reset-size11.size9 {\n  font-size: 0.69453376em;\n}\n.katex .sizing.reset-size11.size10,\n.katex .fontsize-ensurer.reset-size11.size10 {\n  font-size: 0.83360129em;\n}\n.katex .sizing.reset-size11.size11,\n.katex .fontsize-ensurer.reset-size11.size11 {\n  font-size: 1em;\n}\n.katex .delimsizing.size1 {\n  font-family: KaTeX_Size1;\n}\n.katex .delimsizing.size2 {\n  font-family: KaTeX_Size2;\n}\n.katex .delimsizing.size3 {\n  font-family: KaTeX_Size3;\n}\n.katex .delimsizing.size4 {\n  font-family: KaTeX_Size4;\n}\n.katex .delimsizing.mult .delim-size1 > span {\n  font-family: KaTeX_Size1;\n}\n.katex .delimsizing.mult .delim-size4 > span {\n  font-family: KaTeX_Size4;\n}\n.katex .nulldelimiter {\n  display: inline-block;\n  width: 0.12em;\n}\n.katex .delimcenter {\n  position: relative;\n}\n.katex .op-symbol {\n  position: relative;\n}\n.katex .op-symbol.small-op {\n  font-family: KaTeX_Size1;\n}\n.katex .op-symbol.large-op {\n  font-family: KaTeX_Size2;\n}\n.katex .op-limits > .vlist-t {\n  text-align: center;\n}\n.katex .accent > .vlist-t {\n  text-align: center;\n}\n.katex .accent .accent-body {\n  width: 0;\n  position: relative;\n}\n.katex .overlay {\n  display: block;\n}\n.katex .mtable .vertical-separator {\n  display: inline-block;\n  margin: 0 -0.125em;\n  width: 0.25em;\n}\n.katex .mtable .arraycolsep {\n  display: inline-block;\n}\n.katex .mtable .col-align-c > .vlist-t {\n  text-align: center;\n}\n.katex .mtable .col-align-l > .vlist-t {\n  text-align: left;\n}\n.katex .mtable .col-align-r > .vlist-t {\n  text-align: right;\n}\n.katex .svg-align {\n  text-align: left;\n}\n.katex svg {\n  display: block;\n  position: absolute;\n  width: 100%;\n  fill: currentColor;\n  stroke: currentColor;\n  fill-rule: nonzero;\n  fill-opacity: 1;\n  stroke-width: 1;\n  stroke-linecap: butt;\n  stroke-linejoin: miter;\n  stroke-miterlimit: 4;\n  stroke-dasharray: none;\n  stroke-dashoffset: 0;\n  stroke-opacity: 1;\n}\n.katex svg path {\n  stroke: none;\n}\n.katex .vertical-separator svg {\n  width: 0.25em;\n}\n.katex .stretchy {\n  width: 100%;\n  display: block;\n}\n.katex .stretchy:before,\n.katex .stretchy:after {\n  content: \"\";\n}\n.katex .hide-tail {\n  width: 100%;\n  position: relative;\n  overflow: hidden;\n}\n.katex .halfarrow-left {\n  position: absolute;\n  left: 0;\n  width: 50.1%;\n  overflow: hidden;\n}\n.katex .halfarrow-right {\n  position: absolute;\n  right: 0;\n  width: 50%;\n  overflow: hidden;\n}\n.katex .brace-left {\n  position: absolute;\n  left: 0;\n  width: 25.1%;\n  overflow: hidden;\n}\n.katex .brace-center {\n  position: absolute;\n  left: 25%;\n  width: 50%;\n  overflow: hidden;\n}\n.katex .brace-right {\n  position: absolute;\n  right: 0;\n  width: 25.1%;\n  overflow: hidden;\n}\n.katex .x-arrow-pad {\n  padding: 0 0.5em;\n}\n.katex .x-arrow,\n.katex .mover,\n.katex .munder {\n  text-align: center;\n}\n.katex .boxpad {\n  padding: 0 0.3em 0 0.3em;\n}\n.katex .fbox {\n  box-sizing: border-box;\n  border: 0.04em solid black;\n}\n.katex .fcolorbox {\n  box-sizing: border-box;\n  border: 0.04em solid;\n}\n.katex .cancel-pad {\n  padding: 0 0.2em 0 0.2em;\n}\n.katex .mord + .cancel-lap,\n.katex .mbin + .cancel-lap {\n  margin-left: -0.2em;\n}\n.katex .cancel-lap + .mord,\n.katex .cancel-lap + .mbin,\n.katex .cancel-lap + .msupsub {\n  margin-left: -0.2em;\n}\n.katex .sout {\n  border-bottom-style: solid;\n  border-bottom-width: 0.08em;\n}\n", ""]);

// exports


/***/ }),
/* 9 */
/***/ (function(module, exports) {

module.exports = function escape(url) {
    if (typeof url !== 'string') {
        return url
    }
    // If url is already wrapped in quotes, remove them
    if (/^['"].*['"]$/.test(url)) {
        url = url.slice(1, -1);
    }
    // Should url be wrapped?
    // See https://drafts.csswg.org/css-values-3/#urls
    if (/["'() \t\n]/.test(url)) {
        return '"' + url.replace(/"/g, '\\"').replace(/\n/g, '\\n') + '"'
    }

    return url
}


/***/ }),
/* 10 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_AMS-Regular.woff2?3d8245dcb4489694a6a263b05c1cca01";

/***/ }),
/* 11 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_AMS-Regular.woff?ac1d46d953d403677171697581a284d2";

/***/ }),
/* 12 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_AMS-Regular.ttf?c67be87adba7d31c013be127b936233d";

/***/ }),
/* 13 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Caligraphic-Bold.woff2?970d3e76493b82fccf21ad5888ddee77";

/***/ }),
/* 14 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Caligraphic-Bold.woff?445f96a387df0d13ded71f27c608516d";

/***/ }),
/* 15 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Caligraphic-Bold.ttf?3f61a84d76e80396489d32bc9dd8d444";

/***/ }),
/* 16 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Caligraphic-Regular.woff2?0ef0f2e356a2e1c457b6585d34edae77";

/***/ }),
/* 17 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Caligraphic-Regular.woff?74f6918c7d2b768ffd32048102bc0172";

/***/ }),
/* 18 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Caligraphic-Regular.ttf?c3bc8fcec0e85a50cabf71e4e8074991";

/***/ }),
/* 19 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Fraktur-Bold.woff2?950649ba5e5cfd37cdad74095411d350";

/***/ }),
/* 20 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Fraktur-Bold.woff?1aca7ef7f976fb54a207ffc8aa180e38";

/***/ }),
/* 21 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Fraktur-Bold.ttf?e11e6bf02fc41279a540d3655abf3b07";

/***/ }),
/* 22 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Fraktur-Regular.woff2?135ccd74931753318f6f52f3fce19018";

/***/ }),
/* 23 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Fraktur-Regular.woff?c5b430bfcb3e5423b77401afcdb69b66";

/***/ }),
/* 24 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Fraktur-Regular.ttf?a9509497466d16e6b7265a223ea39093";

/***/ }),
/* 25 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Main-Bold.woff2?82bce349c8ff0927380b5177e9258ad0";

/***/ }),
/* 26 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Main-Bold.woff?29d4b276c622ff1d1376d9afcbd2f25a";

/***/ }),
/* 27 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Main-Bold.ttf?db576c7d5d6eb6ebefc9b334f4140d42";

/***/ }),
/* 28 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Main-BoldItalic.woff2?13900e552dbf51aead905f897043226c";

/***/ }),
/* 29 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Main-BoldItalic.woff?3f3be43bbdd58a4a86c8eddd3d39a26a";

/***/ }),
/* 30 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Main-BoldItalic.ttf?cec43f729e0e6810c3433f6020ec676e";

/***/ }),
/* 31 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Main-Italic.woff2?b488f60f79c9d1f533def644385d8219";

/***/ }),
/* 32 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Main-Italic.woff?f8a754484e776e94b6116dfbcb1761a7";

/***/ }),
/* 33 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Main-Italic.ttf?85b3329fdb41fadfb0fdc01ec5290b2f";

/***/ }),
/* 34 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Main-Regular.woff2?999cd6bac6e8118065ac500d4f133c63";

/***/ }),
/* 35 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Main-Regular.woff?485696cf1d821baac6ae06c42e09faef";

/***/ }),
/* 36 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Main-Regular.ttf?f4fe6dca4a5f5520ae0caab31848aae3";

/***/ }),
/* 37 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Math-Italic.woff2?7a31741a44e58952cb4b8a763c206fcd";

/***/ }),
/* 38 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Math-Italic.woff?8eb56b3ff5b141cd3732a24e65c2b339";

/***/ }),
/* 39 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Math-Italic.ttf?10740d7488d690a743339650103f1cb3";

/***/ }),
/* 40 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_SansSerif-Bold.woff2?0d0f967a34b828083a11bb264e578c32";

/***/ }),
/* 41 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_SansSerif-Bold.woff?99bd87431ac7ef8f27591f72d7812509";

/***/ }),
/* 42 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_SansSerif-Bold.ttf?771a1f955e7ff2106402b58d1e8a1ad2";

/***/ }),
/* 43 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_SansSerif-Italic.woff2?b1ac79b9b69c954bbeb8fc02500832e0";

/***/ }),
/* 44 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_SansSerif-Italic.woff?9c3f03ec809c6f298d716cbda8260fcb";

/***/ }),
/* 45 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_SansSerif-Italic.ttf?cb729d47fbd26ff0ae62bb840085e5b5";

/***/ }),
/* 46 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_SansSerif-Regular.woff2?23f1730fda70ec6ec6aa789979890cba";

/***/ }),
/* 47 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_SansSerif-Regular.woff?fc90839714a9e877682309c3daf7493e";

/***/ }),
/* 48 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_SansSerif-Regular.ttf?43b0cb92c2405f175772931ccd68002e";

/***/ }),
/* 49 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Script-Regular.woff2?ae1fad1f4d1c227c9d567da8ea9f988c";

/***/ }),
/* 50 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Script-Regular.woff?60febfa114c5e32f0ce73050476aa39f";

/***/ }),
/* 51 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Script-Regular.ttf?e9169ca7b32608b6235ec9ffff742a71";

/***/ }),
/* 52 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Size1-Regular.woff2?fda0824e8cfb676664ad063c367cfce3";

/***/ }),
/* 53 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Size1-Regular.woff?1d6cf6b89f694dc76e13ecc4e8214dd2";

/***/ }),
/* 54 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Size1-Regular.ttf?f2e296ef7cef0f6f16912e5d171929b2";

/***/ }),
/* 55 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Size2-Regular.woff2?a7a0eb93afb696cb9b4efb87e6cd45a3";

/***/ }),
/* 56 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Size2-Regular.woff?f48e83b7bcaa68a86a78e8edfee1a04e";

/***/ }),
/* 57 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Size2-Regular.ttf?0767ede65042583bcc06f09055f4d7ca";

/***/ }),
/* 58 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Size3-Regular.woff2?397bd6872be7adff4f000c49fa85a9c3";

/***/ }),
/* 59 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Size3-Regular.woff?0387ab387ace32f15b43c3a4a39f187f";

/***/ }),
/* 60 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Size3-Regular.ttf?882f723427a0986846fe786bd8d6f6c7";

/***/ }),
/* 61 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Size4-Regular.woff2?3c14ed11dd8eea981c93bf283193856d";

/***/ }),
/* 62 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Size4-Regular.woff?726db3bd435c8a6e45e44741dae67263";

/***/ }),
/* 63 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Size4-Regular.ttf?27901d5ce93d8971c416d9123fedb911";

/***/ }),
/* 64 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Typewriter-Regular.woff2?895baecfbf4ef298fdfe943fb117c15e";

/***/ }),
/* 65 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Typewriter-Regular.woff?696705de367f02534e8abb38b55af067";

/***/ }),
/* 66 */
/***/ (function(module, exports) {

module.exports = "/fonts/vendor/katex/dist/KaTeX_Typewriter-Regular.ttf?6374f53e08c02d7b9ec594e2fb0c5fe8";

/***/ }),
/* 67 */
/***/ (function(module, exports, __webpack_require__) {

/*
	MIT License http://www.opensource.org/licenses/mit-license.php
	Author Tobias Koppers @sokra
*/

var stylesInDom = {};

var	memoize = function (fn) {
	var memo;

	return function () {
		if (typeof memo === "undefined") memo = fn.apply(this, arguments);
		return memo;
	};
};

var isOldIE = memoize(function () {
	// Test for IE <= 9 as proposed by Browserhacks
	// @see http://browserhacks.com/#hack-e71d8692f65334173fee715c222cb805
	// Tests for existence of standard globals is to allow style-loader
	// to operate correctly into non-standard environments
	// @see https://github.com/webpack-contrib/style-loader/issues/177
	return window && document && document.all && !window.atob;
});

var getElement = (function (fn) {
	var memo = {};

	return function(selector) {
		if (typeof memo[selector] === "undefined") {
			memo[selector] = fn.call(this, selector);
		}

		return memo[selector]
	};
})(function (target) {
	return document.querySelector(target)
});

var singleton = null;
var	singletonCounter = 0;
var	stylesInsertedAtTop = [];

var	fixUrls = __webpack_require__(68);

module.exports = function(list, options) {
	if (typeof DEBUG !== "undefined" && DEBUG) {
		if (typeof document !== "object") throw new Error("The style-loader cannot be used in a non-browser environment");
	}

	options = options || {};

	options.attrs = typeof options.attrs === "object" ? options.attrs : {};

	// Force single-tag solution on IE6-9, which has a hard limit on the # of <style>
	// tags it will allow on a page
	if (!options.singleton) options.singleton = isOldIE();

	// By default, add <style> tags to the <head> element
	if (!options.insertInto) options.insertInto = "head";

	// By default, add <style> tags to the bottom of the target
	if (!options.insertAt) options.insertAt = "bottom";

	var styles = listToStyles(list, options);

	addStylesToDom(styles, options);

	return function update (newList) {
		var mayRemove = [];

		for (var i = 0; i < styles.length; i++) {
			var item = styles[i];
			var domStyle = stylesInDom[item.id];

			domStyle.refs--;
			mayRemove.push(domStyle);
		}

		if(newList) {
			var newStyles = listToStyles(newList, options);
			addStylesToDom(newStyles, options);
		}

		for (var i = 0; i < mayRemove.length; i++) {
			var domStyle = mayRemove[i];

			if(domStyle.refs === 0) {
				for (var j = 0; j < domStyle.parts.length; j++) domStyle.parts[j]();

				delete stylesInDom[domStyle.id];
			}
		}
	};
};

function addStylesToDom (styles, options) {
	for (var i = 0; i < styles.length; i++) {
		var item = styles[i];
		var domStyle = stylesInDom[item.id];

		if(domStyle) {
			domStyle.refs++;

			for(var j = 0; j < domStyle.parts.length; j++) {
				domStyle.parts[j](item.parts[j]);
			}

			for(; j < item.parts.length; j++) {
				domStyle.parts.push(addStyle(item.parts[j], options));
			}
		} else {
			var parts = [];

			for(var j = 0; j < item.parts.length; j++) {
				parts.push(addStyle(item.parts[j], options));
			}

			stylesInDom[item.id] = {id: item.id, refs: 1, parts: parts};
		}
	}
}

function listToStyles (list, options) {
	var styles = [];
	var newStyles = {};

	for (var i = 0; i < list.length; i++) {
		var item = list[i];
		var id = options.base ? item[0] + options.base : item[0];
		var css = item[1];
		var media = item[2];
		var sourceMap = item[3];
		var part = {css: css, media: media, sourceMap: sourceMap};

		if(!newStyles[id]) styles.push(newStyles[id] = {id: id, parts: [part]});
		else newStyles[id].parts.push(part);
	}

	return styles;
}

function insertStyleElement (options, style) {
	var target = getElement(options.insertInto)

	if (!target) {
		throw new Error("Couldn't find a style target. This probably means that the value for the 'insertInto' parameter is invalid.");
	}

	var lastStyleElementInsertedAtTop = stylesInsertedAtTop[stylesInsertedAtTop.length - 1];

	if (options.insertAt === "top") {
		if (!lastStyleElementInsertedAtTop) {
			target.insertBefore(style, target.firstChild);
		} else if (lastStyleElementInsertedAtTop.nextSibling) {
			target.insertBefore(style, lastStyleElementInsertedAtTop.nextSibling);
		} else {
			target.appendChild(style);
		}
		stylesInsertedAtTop.push(style);
	} else if (options.insertAt === "bottom") {
		target.appendChild(style);
	} else {
		throw new Error("Invalid value for parameter 'insertAt'. Must be 'top' or 'bottom'.");
	}
}

function removeStyleElement (style) {
	if (style.parentNode === null) return false;
	style.parentNode.removeChild(style);

	var idx = stylesInsertedAtTop.indexOf(style);
	if(idx >= 0) {
		stylesInsertedAtTop.splice(idx, 1);
	}
}

function createStyleElement (options) {
	var style = document.createElement("style");

	options.attrs.type = "text/css";

	addAttrs(style, options.attrs);
	insertStyleElement(options, style);

	return style;
}

function createLinkElement (options) {
	var link = document.createElement("link");

	options.attrs.type = "text/css";
	options.attrs.rel = "stylesheet";

	addAttrs(link, options.attrs);
	insertStyleElement(options, link);

	return link;
}

function addAttrs (el, attrs) {
	Object.keys(attrs).forEach(function (key) {
		el.setAttribute(key, attrs[key]);
	});
}

function addStyle (obj, options) {
	var style, update, remove, result;

	// If a transform function was defined, run it on the css
	if (options.transform && obj.css) {
	    result = options.transform(obj.css);

	    if (result) {
	    	// If transform returns a value, use that instead of the original css.
	    	// This allows running runtime transformations on the css.
	    	obj.css = result;
	    } else {
	    	// If the transform function returns a falsy value, don't add this css.
	    	// This allows conditional loading of css
	    	return function() {
	    		// noop
	    	};
	    }
	}

	if (options.singleton) {
		var styleIndex = singletonCounter++;

		style = singleton || (singleton = createStyleElement(options));

		update = applyToSingletonTag.bind(null, style, styleIndex, false);
		remove = applyToSingletonTag.bind(null, style, styleIndex, true);

	} else if (
		obj.sourceMap &&
		typeof URL === "function" &&
		typeof URL.createObjectURL === "function" &&
		typeof URL.revokeObjectURL === "function" &&
		typeof Blob === "function" &&
		typeof btoa === "function"
	) {
		style = createLinkElement(options);
		update = updateLink.bind(null, style, options);
		remove = function () {
			removeStyleElement(style);

			if(style.href) URL.revokeObjectURL(style.href);
		};
	} else {
		style = createStyleElement(options);
		update = applyToTag.bind(null, style);
		remove = function () {
			removeStyleElement(style);
		};
	}

	update(obj);

	return function updateStyle (newObj) {
		if (newObj) {
			if (
				newObj.css === obj.css &&
				newObj.media === obj.media &&
				newObj.sourceMap === obj.sourceMap
			) {
				return;
			}

			update(obj = newObj);
		} else {
			remove();
		}
	};
}

var replaceText = (function () {
	var textStore = [];

	return function (index, replacement) {
		textStore[index] = replacement;

		return textStore.filter(Boolean).join('\n');
	};
})();

function applyToSingletonTag (style, index, remove, obj) {
	var css = remove ? "" : obj.css;

	if (style.styleSheet) {
		style.styleSheet.cssText = replaceText(index, css);
	} else {
		var cssNode = document.createTextNode(css);
		var childNodes = style.childNodes;

		if (childNodes[index]) style.removeChild(childNodes[index]);

		if (childNodes.length) {
			style.insertBefore(cssNode, childNodes[index]);
		} else {
			style.appendChild(cssNode);
		}
	}
}

function applyToTag (style, obj) {
	var css = obj.css;
	var media = obj.media;

	if(media) {
		style.setAttribute("media", media)
	}

	if(style.styleSheet) {
		style.styleSheet.cssText = css;
	} else {
		while(style.firstChild) {
			style.removeChild(style.firstChild);
		}

		style.appendChild(document.createTextNode(css));
	}
}

function updateLink (link, options, obj) {
	var css = obj.css;
	var sourceMap = obj.sourceMap;

	/*
		If convertToAbsoluteUrls isn't defined, but sourcemaps are enabled
		and there is no publicPath defined then lets turn convertToAbsoluteUrls
		on by default.  Otherwise default to the convertToAbsoluteUrls option
		directly
	*/
	var autoFixUrls = options.convertToAbsoluteUrls === undefined && sourceMap;

	if (options.convertToAbsoluteUrls || autoFixUrls) {
		css = fixUrls(css);
	}

	if (sourceMap) {
		// http://stackoverflow.com/a/26603875
		css += "\n/*# sourceMappingURL=data:application/json;base64," + btoa(unescape(encodeURIComponent(JSON.stringify(sourceMap)))) + " */";
	}

	var blob = new Blob([css], { type: "text/css" });

	var oldSrc = link.href;

	link.href = URL.createObjectURL(blob);

	if(oldSrc) URL.revokeObjectURL(oldSrc);
}


/***/ }),
/* 68 */
/***/ (function(module, exports) {


/**
 * When source maps are enabled, `style-loader` uses a link element with a data-uri to
 * embed the css on the page. This breaks all relative urls because now they are relative to a
 * bundle instead of the current page.
 *
 * One solution is to only use full urls, but that may be impossible.
 *
 * Instead, this function "fixes" the relative urls to be absolute according to the current page location.
 *
 * A rudimentary test suite is located at `test/fixUrls.js` and can be run via the `npm test` command.
 *
 */

module.exports = function (css) {
  // get current location
  var location = typeof window !== "undefined" && window.location;

  if (!location) {
    throw new Error("fixUrls requires window.location");
  }

	// blank or null?
	if (!css || typeof css !== "string") {
	  return css;
  }

  var baseUrl = location.protocol + "//" + location.host;
  var currentDir = baseUrl + location.pathname.replace(/\/[^\/]*$/, "/");

	// convert each url(...)
	/*
	This regular expression is just a way to recursively match brackets within
	a string.

	 /url\s*\(  = Match on the word "url" with any whitespace after it and then a parens
	   (  = Start a capturing group
	     (?:  = Start a non-capturing group
	         [^)(]  = Match anything that isn't a parentheses
	         |  = OR
	         \(  = Match a start parentheses
	             (?:  = Start another non-capturing groups
	                 [^)(]+  = Match anything that isn't a parentheses
	                 |  = OR
	                 \(  = Match a start parentheses
	                     [^)(]*  = Match anything that isn't a parentheses
	                 \)  = Match a end parentheses
	             )  = End Group
              *\) = Match anything and then a close parens
          )  = Close non-capturing group
          *  = Match anything
       )  = Close capturing group
	 \)  = Match a close parens

	 /gi  = Get all matches, not the first.  Be case insensitive.
	 */
	var fixedCss = css.replace(/url\s*\(((?:[^)(]|\((?:[^)(]+|\([^)(]*\))*\))*)\)/gi, function(fullMatch, origUrl) {
		// strip quotes (if they exist)
		var unquotedOrigUrl = origUrl
			.trim()
			.replace(/^"(.*)"$/, function(o, $1){ return $1; })
			.replace(/^'(.*)'$/, function(o, $1){ return $1; });

		// already a full url? no change
		if (/^(#|data:|http:\/\/|https:\/\/|file:\/\/\/)/i.test(unquotedOrigUrl)) {
		  return fullMatch;
		}

		// convert the url to a full url
		var newUrl;

		if (unquotedOrigUrl.indexOf("//") === 0) {
		  	//TODO: should we add protocol?
			newUrl = unquotedOrigUrl;
		} else if (unquotedOrigUrl.indexOf("/") === 0) {
			// path should be relative to the base url
			newUrl = baseUrl + unquotedOrigUrl; // already starts with '/'
		} else {
			// path should be relative to current directory
			newUrl = currentDir + unquotedOrigUrl.replace(/^\.\//, ""); // Strip leading './'
		}

		// send back the fixed url(...)
		return "url(" + JSON.stringify(newUrl) + ")";
	});

	// send back the fixed css
	return fixedCss;
};


/***/ })
/******/ ]);