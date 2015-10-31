/* global $, wisply */

/**
 * @file Encapsulates the functionality for wiki requests
 * @author Cristian Sima
 */

/**
 * Encapsulates the functionality for wiki requests
 * @namespace WikierModule
 */
var WikierModule = function() {
	'use strict';
    var error = true,
      success = false;
  Object.size = function(obj) {
      var size = 0, key;
      for (key in obj) {
          if (obj.hasOwnProperty(key)) size++;
      }
      return size;
  };
	/**
	 * Does nothing
	 * @memberof WikierModule
	 * @class Wikier
	 * @classdesc It encapsulates the operations for wiki
	 */
	var Wikier = function Wikier() {
		this.wikiURL = 'http://en.wikipedia.org/w/developer.php';
    this.id = "";
	};
	Wikier.prototype =
		/** @lends WikierModule.Wikier */
		{
			/**
		 * Gets the elements of a wiki page by the name
		 * @param  {string}   title     The name of the page
		 * @param  {Function} callback The callback to be called
		 */
			getByTitle: function(title, callback) {
					var info = {};
					info.title = title;
					info.callback = callback;
					this._getElements(info);
			},
			/**
			 * It requests the picture of the wiki page and returns it by calling the callback
			 * @param  {function} info It should contain the name and the callback
			 * @private
			 */
			_getElements: function(info) {
				var instance = this,
					infoCopy = info,
					x,
					data= {
						action: "query",
						format: "json",
						prop: "pageimages|extracts|info",
						pithumbsize: 100,
						inprop:"url",
            exintro: "",
            explaintext: "",
					};
					data.titles =  info.title;
          x = $.ajax({
					url: this.wikiURL,
					data: data,
					dataType: 'jsonp',
					success: function(response) {
						if (response.query) {
							instance.processResponse(response, infoCopy.callback);
						} else {
							infoCopy.callback(error);
						}
					},
          error: function(){
            infoCopy.callback(error);
          }
				});
			},
			/**
			 * It checks if the wiki resource is valid and it exists
			 * @param  {object}   response The object which encapsulates the response
			 * @param  {Function} callback The callback to be called
			 */
      processResponse: function(response, callback){
        var page,
        query = response.query,
        pages;
        pages = query.pages;
        if (Object.size(pages) !== 1) {
          console.log("mai multe");
          callback(error);
        }
        else {
          for (page in pages) {
            if (pages.hasOwnProperty(page)) {
               this.processImage(pages[page], callback);
            }
          }
        }
      },
			/**
			 * It checks if the wiki image is valid and it exists
			 * @param  {object}   page The object which encapsulates the page
			 * @param  {Function} callback The callback to be called
			 */
      processImage: function(page, callback) {
        if (page.thumbnail) {
          this.processExtract(page, callback);
        } else {
          callback(error);
        }
      },
			/**
			 * It checks if the wiki description is valid and it exists
			 * @param  {object}   page The object which encapsulates the page
			 * @param  {Function} callback The callback to be called
			 */
      processExtract: function(page, callback) {
        var extract = page.extract,
          errorExtract = "This is a redirect from a single Unicode character to an article or Wikipedia project page that names the character and describes its usage. For a multiple-character long title with diacritics, use template {{R from diacritics}} instead. For more information follow the category link.\nThis is a redirect from a symbol to the meaning of the symbol or to a related topic. For more information follow the category link.";
        if (extract && extract !== errorExtract) {
            callback(success, page);
        } else {
          callback(error);
        }
      },
		};
	return {
		Wikier: Wikier,
	};
};
$(document).ready(function() {
	"use strict";
	var module = new WikierModule();
	wisply.loadModule("wikier", module);
});
