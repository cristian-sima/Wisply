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
	/**
	 * Does nothing
	 * @memberof WikierModule
	 * @class Wikier
	 * @classdesc It encapsulates the operations for wiki
	 */
	var Wikier = function Wikier() {
		this.wikiURL = 'http://en.wikipedia.org/w/api.php';
		this.subject = "";
	};
	Wikier.prototype =
		/** @lends WikierModule.Wikier */
		{
			/**
			 * It changes the subject of the wiki
			 * @param  {string} subject The new subject
			 */
			changeSubject: function(subject) {
				this.subject = subject;
			},
			/**
			 * It requests the picture of the wiki page and returns it by calling the callback
			 * @param  {function} callback It is called when the picture is received
			 */
			getPicture: function(callback) {
				var error = true,
					success = false;
				$.ajax({
					url: this.wikiURL,
					data: {
						action: "query",
						titles: this.subject,
						prop: "pageimages",
						format: "json",
						pithumbsize: 100,
					},
					dataType: 'jsonp',
					success: function(response) {
						var page,
							thumbnail,
							query = response.query,
							pages;
						if (query) {
							pages = query.pages;
							for (page in pages) {
								if (pages.hasOwnProperty(page)) {
									thumbnail = pages[page].thumbnail;
									if (thumbnail) {
										callback(success, thumbnail);
									} else {
										callback(error);
									}
								}
							}
						} else {
							callback(error);
						}
					},
          error: function(){
            callback(error);
          }
				});
			},
			/**
			 * It gets the short description and returns it by calling the callback
			 * @param  {function} callback It is called when the description is received
			 */
			getDescription: function(callback) {
				var error = true,
					success = false;
				$.ajax({
					url: this.wikiURL,
					data: {
						action: "query",
						titles: this.subject,
						prop: "extracts",
						format: "json",
            exintro: "",
            explaintext: "",
					},
					dataType: 'jsonp',
					success: function(response) {
						var page,
							extract,
							query = response.query,
							pages,
              errorExtract = "This is a redirect from a single Unicode character to an article or Wikipedia project page that names the character and describes its usage. For a multiple-character long title with diacritics, use template {{R from diacritics}} instead. For more information follow the category link.\nThis is a redirect from a symbol to the meaning of the symbol or to a related topic. For more information follow the category link.";
						if (query) {
							pages = query.pages;
							for (page in pages) {
								if (pages.hasOwnProperty(page)) {
									extract = pages[page].extract;
									if (extract) {
                    if(extract === errorExtract) {
                      callback(error);
                    } else {
										  callback(success, extract);
                    }
									} else {
										callback(error);
									}
								}
							}
						} else {
							callback(error);
						}
					},
          error: function(){
            callback(error);
          },
				});
			}
		};
	return {
		Wikier: Wikier,
	};
};
$(document).ready(function() {
	"use strict";
	var module = new WikierModule();
	wisply.wikierModule = module;
});
