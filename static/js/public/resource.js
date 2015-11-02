/* global $, wisply */
/**
 * @file Encapsulates the functionality for displaying a public resource
 * @author Cristian Sima
 */
/**
 * Requires nothing
 * @namespace PublicResourceModule
 */
var PublicResourceModule = function() {
	'use strict';
	var page;
	/**
	 * Encapsulates the functionality for public resources
	 * @class Page
	 * @memberof PublicResourceModule
	 * @classdesc It represents the page
	 * @param {object} o The object with the repository and resource objects
	 */
	var Page = function Page(o) {
		this.similarCategory = new DIV(this, {
			selector: "#div-same-collection",
			url: "/developer/api/repository/resources/" + o.repository.id + "/get/0/10?collection=&format=html",
		});
		this.repository = o.repository;
		this.resource = o.resource;
	};
	Page.prototype =
		/** @lends PublicResourceModule.Page */
		{
			/**
			 * It inits the divs
			 */
			init: function() {
				this.similarCategory.init();
			},
		};
	/**
	 * @class DIV
	 * @memberof PublicResourceModule
	 * @classdesc It represents a div with similar
	 * @param {PublicResourceModule.Page} page The reference to the page object
	 * @param {string} o The options for the div. It must contain: selector, URL
	 */
	var DIV = function DIV(page, o) {
		this.element = $(o.selector);
		this.page = page;
		this.url = o.url;
	};
	DIV.prototype =
		/** @lends PublicResourceModule.DIV */
		{
			/**
			 * It shows the loading
			 */
		 init: function() {
				this._showLoading();
				this.load();
			},
			/**
			 * It shows the loading image
			 */
			_showLoading: function() {
				//src="/repository/{{ .repository.ID }}/resource/{{ .resource.ID }}/content"
				var html = "<div style='text-align:center;margin-top:50px; margin-bottom:50px;' >" + wisply.getLoadingImage("medium") + "<br />Loading content...</div>";
				this._setContent(html);
			},
			/**
			 * It changes the content of the div
			 * @param  {string} content The new content to be shown
			 */
			_setContent: function(content) {
				this.element.html(content);
				wisply.activateTooltip();
			},
			/**
			 * It loads the URL and shows the HTML in the div
			 */
			load: function () {
				var instance = this;
				$.ajax({
					url: this.url,
					data: {
						"format": "html",
					},
					success: function(html) {
						var seeMore = "<br /><a href='/repository/" + instance.page.repository.id + "#list|10-15*'> See more resources from EdShare </a>";
						instance._setContent(html + seeMore);
					}
				});
			},
		};
	/**
	 * Creates the page object and stores it in the page variable
	 * @param  {object} object The object which contains the repository and the
	 * resource
	 * @return {PublicResourceModule.Page} The page
	 */
	function init(object) {
		page = new Page(object);
		page.init();
		return page;
	}
	return {
		init: init,
		Page: Page
	};
};
$(document).ready(function() {
	"use strict";
	wisply.loadModule("public-resource", new PublicResourceModule());
});
