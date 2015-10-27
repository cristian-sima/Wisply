/* global $, wisply, Handlebars */
/**
 * @file Encapsulates the functionality for managing captcha images
 * @author Cristian Sima
 */
/**
 * @namespace CaptchaModule
 */
var CaptchaModule = function() {
	'use strict';
	var settings = {
			noEscape: true,
		},
		pathToFolder = "/captcha/";
	/**
	 * It loads the image and it calls the callback when image is loaded
	 * @param  {string}   path     The path of the image
	 * @param  {Function} callback The function which is called after the image has loaded
	 */
	function loadImage(path, callback) {
		$('<img class="captcha-image thumbnail" alt="Catcha Image" src="' + path + '">').load(function() {
			callback(this);
		});
	}
	/**
	 * The constructor does nothing
	 * @class Captcha
	 * @memberof CaptchaModule
	 * @classdesc The Manager for Captcha images
	 * @param {object} o The object which contains the options. It must have: ID and name
	 */
	function Captcha(o) {
		this.ID = o.ID;
		this.name = "captcha-" + o.name;
		this.element = o.element;
		this.reloadParameter = "";
	}
	Captcha.prototype =
		/** @lends Captcha.CaptchaModule */
		{
			/**
			 * It shows the loading icon
			 */
			_showLoading: function() {
				var html = "<div class='text-center' style='padding-top:100px'>{{ image }}</div>",
					template = Handlebars.compile(html, settings),
					data = {
						image: wisply.getLoadingImage("medium"),
					},
					content = template(data);
				this.element.html(content);
			},
			/**
			 * It loads the image from the server and displays the DIV
			 */
			show: function() {
				this._loadImage();
			},
			/**
			 * It sets the reloadParameter as ?reload=timestamp and loads the
			 * image from the server
			 */
			reload: function() {
				/**
				 * It returns the current timestamp
				 * @return {string} The current timestamp
				 */
				function getCurrentTime() {
					return new Date().getTime();
				}
				this.reloadParameter = "?reload=" + getCurrentTime();
				this._loadImage();
			},
			/**
			 * It shows the loading image, and calls the loadImage function with
			 * _fired_imageLoaded as the callback
			 */
			_loadImage: function() {
				this._showLoading();
				var instance = this,
					path = this._getPath();
				loadImage(path, function(image) {
					instance._fired_imageLoaded(image);
				});
			},
			/**
			 * It returns the path to the image
			 * @return {string} The path to the image
			 */
			_getPath: function() {
				return pathToFolder + this.ID + ".png" + this.reloadParameter;
			},
			/**
			 * It is called when the image has been loaded. It creates the DIV and
			 * shows it
			 * @param  {object} image The DOM image element
			 */
			_fired_imageLoaded: function(image) {
				/**
				 * It returns the content of the div
				 * @param  {object} image The DOM object of the image
				 * @return {string} The string for the DIV
				 */
				function getContent(image) {
					var html = "<input type='hidden' value='{{ id }}' name='{{ name }}' />" + "Type the numbers which appear in the next image. <br />" + "Can't read it? Try a <a href='#' id='{{ name }}-reload' data-target='login-form-captcha'> different image</a> or an <a href='#' id='{{ name }}-audio'>audio captcha</a>.<br />" + "{{ image }}" + "<a href='#' class='info-captcha' data-toggle='tooltip' title='This image is intended to distinguish human from machine input. Typically, it is a way of thwarting spam and automated extraction of data from websites.' ><span class='glyphicon glyphicon-question-sign'></span> What's this?</a>",
						values = {
							image: image.outerHTML,
							id: instance.ID,
							name: instance.name,
						},
						template = Handlebars.compile(html, settings);
					return template(values);
				}
				var instance = this,
					content = getContent(image);
				this._updateHTML(content);
			},
			/**
			 * It updates the HTML div
			 * @param  {string} inner The HTML to be inserted
			 */
			_updateHTML: function(inner) {
				var html = "<div class='well' >" + inner + "</div>";
				this.element.html(html);
				this._activateListeners();
			},
			/**
			 * It activates the listeners
			 */
			_activateListeners: function() {
				var instance = this;
				$("#" + this.name + "-reload").click(function(event) {
					event.preventDefault();
					instance.reload();
				});
				$("#" + this.name + "-audio").click(function(event) {
					event.preventDefault();
					/**
					 * It returns the DIV with the HTML code for the audio player
					 * @return {string} HTML code for the audio player
					 */
					function getMessage() {
						var html = '<div><audio id=audio controls autoplay src="/captcha/{{ id }}.wav" preload=none>         You browser does not support audio.         <a href="/captcha/download/{{ id }}.wav">Download file</a> to play it in the external player. </audio></div>',
							values = {
								id: instance.ID,
							},
							template = Handlebars.compile(html, settings);
						return template(values);
					}
					wisply.message.dialog({
						title: "Playing the audio captcha...",
						message: getMessage(),
						onEscape: true,
					});
				});
				wisply.activateTooltip();
			}
		};
	return {
		Captcha: Captcha
	};
};
$(document).ready(function() {
  "use strict";
	wisply.loadModule("captcha", new CaptchaModule());
});
