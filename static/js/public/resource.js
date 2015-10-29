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
		this.iframe = new IFrame(this);
		this.repository = o.repository;
		this.resource = o.resource;
	};
	Page.prototype =
		/** @lends PublicResourceModule.Page */
		{
			/**
			 * It loads the listeners
			 */
			init: function() {
				this.iframe.init();
			},
		};
	/**
	 * @class IFrame
	 * @memberof PublicResourceModule
	 * @classdesc It represents the IFrame
	 * @param {PublicResourceModule.Page} page The reference to the page object
	 */
	var IFrame = function IFrame(page) {
		this.element = $("#the-iframe");
		this.page = page;
	};
	IFrame.prototype =
		/** @lends PublicResourceModule.IFrame */
		{
			/**
			 * It shows the loading
			 */
			init: function() {
				this._showLoading();
			},
			/**
			 * It shows the loading image
			 */
			_showLoading: function() {
				//src="/repository/{{ .repository.ID }}/resource/{{ .resource.ID }}/content"
				var html = "<div style='text-align:center;margin-top:50px; margin-bottom:50px;' >" + wisply.getLoadingImage("medium") + "<br />Loading content...</div>";
				this.element.contents().find("body").html(html);
			}
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
