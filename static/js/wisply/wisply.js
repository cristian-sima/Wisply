/* global $, bootbox, base64_decode, searchModule */
/**
 * It contains a reference to the Wisply.App object
 * @global
 * @see Wisply.App
 */
var wisply;
/**
 * @file Encapsulates the functionality for all pages.
 * @author Cristian Sima
 */
/**
 * Encapsulates the functionality for all pages.
 * @namespace Wisply
 */
var Wisply = function() {
	'use strict';
	/**
	 * The constructor calls the init method
	 * @class ShortcutManager
	 * @memberof Wisply
	 * @classdesc It manages the operations with the key shortcuts
	 */
	function ShortcutManager() {
		this.memory = [];
		this.defaultShortcuts = [{
			"type": "keydown",
			"key": "Alt+a",
			"callback": function() {
				/**
				 * It loads the accessibility javascript file and shows the bar
				 */
				function createScript() {
					var jf = document.createElement('script');
					jf.src = 'https://core.atbar.org/atbar/en/latest/atbar.min.js';
					jf.type = 'text/javascript';
					jf.id = 'ToolBar';
					document.getElementsByTagName('head')[0].appendChild(jf);
				}
				/**
				 * It moves the Wisply navigation bar 41px down. (it prevents overlaping with the bar). Also, it sets a listener for closing button, such that the bar goes back to original position
				 */
				function moveWisply() {
					$(".navbar-fixed-top").css({
						"top": "41px"
					});
					setTimeout(function() {
						$("#at-btn-atkit-unload").click(function() {
							setTimeout(function() {
								$(".navbar-fixed-top").css({
									"top": "0px"
								});
							}, 100);
						});
					}, 200);
				}
				/**
				 * It loads the bar and moves wisply
				 */
				function showAccessibilityBar() {
					createScript();
					moveWisply();
				}
				showAccessibilityBar();
			},
			description: "Show the accessibility bar",
		}, {
			"type": "keydown",
			"key": "Alt+w",
			"callback": function() {
				wisply.goTo("/");
			},
			description: "Load the home page",
		}, {
			"type": "keydown",
			"key": "Alt+c",
			"callback": function() {
				wisply.goTo("/contact");
			},
			description: "Load the contact page",
		}, {
			"type": "keydown",
			"key": "Alt+k",
			"callback": function() {
				var description = wisply.shortcutManager.getDescription();
				wisply.message.show("Key shortcuts available on this page", description);
				wisply.activateTooltip();
			},
			description: "Show the list of key shortcuts",
		}, {
			"type": "keydown",
			"key": "Ctrl+space",
			"callback": function(event) {
				event.preventDefault();
				wisply.search.object.focus();
			},
			description: "Activates the searching",
		},];
	}
	ShortcutManager.prototype =
		/**
		 * @lends Wisply.ShortcutManager
		 */
		{
			/**
			 * Called when the object is create. It activates the default shortcuts
			 */
			init: function() {
				this.activate(this.defaultShortcuts);
			},
			/**
			 * It activates the shortcuts received as parameters
			 * @param  {array} shortcuts An array with the shortcuts to active. A shortcut has a event type, the shortcut combination of keys and the callback
			 */
			activate: function(shortcuts) {
				/**
				 * It binds an shortcut using the JQUERY function.
				 * @param [object] shortcut The shortcut to be binded. It should have type, key and callback
				 */
				function bind(shortcut) {
					$(document).bind(shortcut.type, shortcut.key, shortcut.callback);
				}
				var shortcut, i;
				for (i = 0; i < shortcuts.length; i++) {
					shortcut = shortcuts[i];
					bind(shortcut);
					this.memory.push(shortcut);
				}
			},
			/**
			 * It returns an HTML description of all the current shortcuts
			 * @return {string} The HTML description of all the shortcuts
			 */
			getDescription: function() {
				/**
				 * It returns the HTML description of the keys
				 * @param  {string} keys The string which holds the keys such that: key1+key2
				 * @return {string} The HTML description of the keys
				 */
				function getKeysHTML(keys) {
					/**
					 * It changes for some keys the word to an icon (e.g. arrows)
					 * @param  {string} key The key as a string
					 * @return {string} The HTML of the key
					 */
					function getKey(key) {
						var htmlKey = "",
							defaultKeys = {
								"UP": "<span class='glyphicon glyphicon-arrow-up'></span>",
								"DOWN": "<span class='glyphicon glyphicon-arrow-down'></span>",
								"LEFT": "<span class='glyphicon glyphicon-arrow-left'></span>",
								"RIGHT": "<span class='glyphicon glyphicon-arrow-right'></span>",
							};
						if (defaultKeys[key]) {
							htmlKey = defaultKeys[key];
						} else {
							htmlKey = key;
						}
						return htmlKey;
					}
					var keysHTML = [], index,
						elements = keys.toUpperCase().split("+");
					for (index = 0; index < elements.length; index++) {
						keysHTML.push("<kbd>" + getKey(elements[index]) + "</kbd>");
					}
					return keysHTML.join(" <span class='text-muted'>+</span> ");
				}
				/**
				 * It returns an HTML description of the shortcut
				 * @param  {object} shortcut The shortcut object
				 * @return {string} A description of the shortcut
				 */
				function describeShortcut(shortcut) {
					/**
					 * If the showWarning is true, it returns a warning icon which informs that the shortcut is overwritting the default functionality of the browser
					 * @param  {bool} showWarning If to show or not the warning
					 * @return {string} The warning icon or empty string
					 */
					function getWarning(showWarning) {
						var text = "";
						if (showWarning) {
							text = "<span class='hidden-xs text-warning glyphicon glyphicon-warning-sign' data-placement='top'  data-toggle='tooltip' data-original-title='This key overwrites the default functionality of your browser'></span>&nbsp;&nbsp;";
						}
						return text;
					}
					var html = "<tr><td>";
					html += getKeysHTML(shortcut.key);
					html += "</td><td>";
					html += "&nbsp;&nbsp;" + shortcut.description + "&nbsp;&nbsp;" + getWarning(shortcut.overwrites);
					html += "</td></tr>";
					return html;
				}
				/**
				 * It returns the HTML code which describes the shortcuts
				 * @param  {array} shortcuts The shortcuts
				 * @return {string} HTML code which describes the shortcuts
				 */
				function describeShortcuts(shortcuts) {
					var shortcut,
						text = "";
					for (i = 0; i < shortcuts.length; i++) {
						shortcut = shortcuts[i];
						text += describeShortcut(shortcut);
					}
					return text;
				}
				var i,
					html = "";
				html = "<table class='table table-hover'><tbody>";
				html += describeShortcuts(this.memory);
				html += "</tbody></table>";
				return html;
			}
		};
	/**
	 * It creates the object
	 * @class Message
	 * @memberof Wisply
	 * @classdesc It uses manages the operating regarding JavaScript messages
	 */
	function Message() {
		this.currentMessage = undefined;
	}
	/**
	 * @memberof Message
	 */
	Message.prototype =
		/**
		 * @lends Wisply.Message
		 */
		{
			/**
			 * It shows a succesful message
			 * @param  {string} message The content of the message to be displayed
			 */
			showSuccess: function(message) {
				this.show("<div class='text-success'>Success</div>", message);
			},
			/**
			 * It shows an error message
			 * @param  {string} message The content of the message to be displayed
			 */
			showError: function(message) {
				this.show("<div class='text-warning'>Sorry</div>", message);
			},
			/**
			 * It shows a message
			 * @param  {string} title   The title of the message
			 * @param  {string} content The content of the message
			 */
			show: function(title, content) {
				this.dialog({
					title: title,
					message: content,
					onEscape: function() {},
				});
			},
			/**
			 * It represents an adapter for the bootbox alert function. It shows an error message
			 * @param  {object} args The arguments for the dialog
			 * @see {@link http://bootboxjs.com/|Bootbox official website}
			 */
			alert: function(args) {
				this.dialog(args);
			},
			/**
			 * It represents an adapter for the bootbox alert function. It shows a dialog message
			 * @param  {object} args The arguments for the dialog
			 * @return {object} The bootbox element which holds the message
			 * @see {@link http://bootboxjs.com/|Bootbox official website}
			 */
			dialog: function(args) {
				this._clear();
				return bootbox.dialog(args);
			},
			/**
			 * It shows a waiting message
			 * @param  {string} title The title of the message
			 * @return {object} The bootbox element which holds the message
			 * @see {@link http://bootboxjs.com/|Bootbox official website}
			 */
			tellToWait: function(title) {
				var msg;
				msg = "<div class='text-center text-muted'> It may take up to a minute. Enjoy a coffee (be aware of sugar) :) <br />" + wisply.getLoadingImage("big") + "</div>";
				this.currentMessage = this.dialog({
					title: title,
					message: msg,
				});
				return this.currentMessage;
			},
			/**
			 * In case there is any window open, it hides it.
			 * It is a private method
			 */
			_clear: function() {
				if(this.currentMessage) {
					this.currentMessage.modal('hide');
				}
			}
		};
	/**
	 * The constructor creates a message and a shortcut objects
	 * @property {Message}  message                    The reference to the message object
	 * @property {ShortcutManager}  shortcutManager    The reference to the shortcut manager
	 * @class App
	 * @memberof Wisply
	 * @classdesc It represents the main object of the website. It stores references to other objects and it provides the main functions
	 */
	var App = function App() {
		/**
		 * @access public
		 */
		this.message = new Message();
		/**
		 * @access public
		 */
		this.shortcutManager = new ShortcutManager();
		this.isSmallSearchDisplayed = false;
	};
	App.prototype =
		/**
		 * @lends Wisply.App
		 */
		{
			/**
			 * It initiates the shorcuts' manager
			 */
			init: function() {
				var instance = this;
				function initSearch() {
						function showSmallSearch(element) {
							var smallSearch = $("#search-small-input"),
							button = $(element);
							$("#full-logo").hide();
							$("#search-small").show();
							button.html("");
							smallSearch.focus();
						}
						$("#show-small-search-button").click(function() {
							if(!instance.isSmallSearchDisplayed) {
								showSmallSearch(this);
							}
							instance.isSmallSearchDisplayed = !instance.isSmallSearchDisplayed;
						});
				}
				this.shortcutManager.init();
				this.solveHashProblem();
				initSearch();
				this.search = new searchModule.Field({
					selector: '.wisply-search-field',
					URL: "/api/search/anything/",
					saveSearches: true,
				});
				this.activateTooltip();
			},
			/**
			 * The menu is fixed and thus when the user jumps to #something it does the menu overlapping the content
			 */
			solveHashProblem: function() {
				/**
				 * It scrolls up 80 px and thus it prevents overlaping the elements
				 * This is caused of the fixed navigation
				 */
				function preventNavOverlap() {
					setTimeout(function() {
						scrollBy(0, -80);
					}, 10);
				}
				var hashTagActive = "";
				$(".scroll").click(function(event) {
					/**
					 * It returns the navigation place for the hash
					 */
					function getDestinationPlace(thisObject) {
						var dest = 0;
						if ($(thisObject.hash).offset().top > $(document).height() - $(window).height()) {
							dest = $(document).height() - $(window).height();
						} else {
							dest = $(thisObject.hash).offset().top;
						}
						return dest;
					}
					/**
					 * It changes the hash
					 * @param  {event} object The event which has been generated
					 */
					function changeHash(object) {
						event.preventDefault();
						var dest = getDestinationPlace(object);
						$('html,body').animate({
							scrollTop: dest - 80
						}, 500, 'swing');
						hashTagActive = object.hash;
						window.location.hash = object.hash;
					}
					if (hashTagActive != this.hash) { //this will prevent if the user click several times the same link to freeze the scroll.
						changeHash(this);
					} else {
						preventNavOverlap();
					}
				});
				if (window.location.hash) {
					preventNavOverlap();
				}
				$(window).on('hashchange', function() {
					preventNavOverlap();
				});
			},
			/**
			 * It preloads the loading image in the background and stores it in the browser cache
			 */
			preloadLoadingImage: function() {
				var img = new Image();
				img.src = "/static/img/wisply/load/medium.gif";
			},
			/**
			 * It executes a JQuery post request, adding to it the xsrf token value
			 * @param  {object} args Same arguments for as for a JQuery AJAX request
			 * @see {@link http://api.jquery.com/jquery.ajax/|JQuery AJAX API}
			 */
			executePostAjax: function(args) {
				if (typeof args.data === 'undefined') {
					args.data = {};
				}
				args.dataType = "text";
				args.method = args.type = "POST";
				args.data._xsrf = this.getXSRF();
				$.ajax(args);
			},
			/**
			 * It activates the bootstrap tooltips
			 */
			activateTooltip: function() {
				$('[data-toggle="tooltip"]').tooltip();
			},
			getXSRF: function() {
				var xsrf = $.cookie("_xsrf"),
					xsrflist = xsrf.split("|"),
					value = base64_decode(xsrflist[0]);
				return value;
			},
			/**
			 * It refreshes the page
			 * @param  {number} delayTime The amount of time in ms to delay the refresh
			 */
			reloadPage: function(delayTime) {
				if (typeof size === 'undefined') {} else {
					if (delayTime === "now") {
						delayTime = 0;
					}
				}
				setTimeout(function() {
					location.reload();
				}, delayTime);
			},
			/**
			 * It transforms a HTML object in the loading icon for Wisply
			 * @param  {string} idElement The id of the element
			 * @param  {string} size      The size of the loading icon. It can be small (for 20px), medium (for 55px) and large (for 110px)
			 */
			showLoading: function(idElement, size) {
				var element = $(idElement),
					HTML = this.getLoadingImage(size);
				element.html(HTML);
			},
			/**
			 * It gets the HTML of the loading image
			 * @param  {string} size      The size of the loading icon. It can be small (for 20px), medium (for 55px) and large (for 110px)
			 * @return {string} The HTML for Wisply loading image
			 */
			getLoadingImage: function(size) {
				/**
				 * It returns the dimension in pixels acording to string type
				 * @param  {string} size The demension of the image. It can be small (for 20px), medium (for 55px) and large (for 110px)
				 * @return {number}      The dimension in pixels
				 */
				function getDimension(size) {
					var sizes = {
						"small": 20,
						"medium": 55,
						"big": 110,
					};
					return sizes[size];
				}
				/**
				 * It returns the HTML code for the loading element
				 * @param  {string} size      The size of the loading icon. It can be small (for 20px), medium (for 55px) and large (for 110px)
			 	 * @param  {number} dimension The size of the image in px
				 * @return {string}           The HTML code for loading element
				 */
				function getHTML(size, dimension) {
					return "<img alt='...' src='/static/img/wisply/load/" + size + ".gif" + "' style='height: " + dimension + "px; width: " + dimension + "px' />";
				}
				var dimension;
				if (typeof size === 'undefined') {
					size = "small";
				}
				dimension = getDimension(size);
				return getHTML(size, dimension);
			},
			/**
			 * It redirects the account to a certain page
			 * @param  {string} address The address of the page
			 */
			goTo: function(address) {
				document.location = address;
			},
			/**
			 * It adds a connection to Wisply
			 * @param {Connection} connection The Connection object
			 */
			addConnection: function(connection) {
				this.connection = connection;
			}
		};
	return {
		App: App,
	};
};
$(document).ready(function() {
	"use strict";
	var module = new Wisply();
	wisply = new module.App();
	wisply.init();
});
