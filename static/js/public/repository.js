/* global $, wisply, server */
/**
 * @file Encapsulates the functionality for displaying the contents of a repository for the public
 * @author Cristian Sima
 */
/**
 * @namespace PublicRepository
 */
var PublicRepository = function() {
	'use strict';
	/**
	 * The constructor sets the default values
	 * @memberof PublicRepository
	 * @class Manager
	 * @classdesc It encapsulets the functionality for the public repository
	 * @param {object} currentRepository Represents the current repository
	 */
	var Manager = function Manager(currentRepository) {
		this.repository = currentRepository;
		this.min = 0;
		this.resourcePerPage = 15;
		// They are used to extract the verbs and parameters from the hash
		this.delimitator = {
			verb: "*",
			parameter: "-",
			insideVerb: "|",
		};
	};
	Manager.prototype =
		/** @lends PublicRepository.Manager */
		{
			/**
			 * It activates the listeners and adds the shortcuts
			 */
			init: function() {
				var instance = this;

				function initHash() {
					$(window).on('hashchange', function() {
						instance.hashChanged();
					});
				}
				/**
				 * It adds the keys' shortcuts
				 */
				function initKeys() {
					var shortcuts = [{
						"type": "keyup",
						"key": "Ctrl+left",
						"callback": function() {
							wisply.publicRepositoryModule.manager.showPrevious();
						}
					}, {
						"type": "keyup",
						"key": "Ctrl+right",
						"callback": function() {
							wisply.publicRepositoryModule.manager.showNext();
						}
					}];
					wisply.shortcutManager.activate(shortcuts);
				}
				initHash();
				initKeys();
        this.hashChanged();
			},
			/**
			 * It loads the resources from server according to the current settings
			 */
			getResources: function() {
				var instance = this;
				this.showLoading();
				$.ajax({
					url: "/api/repository/resources/" + this.repository.id + "/get/" + this.min + "/" + this.resourcePerPage,
					success: function(html) {
						instance.fired_NewResources(html);
					}
				});
			},
			/**
			 * It shows the wisply loading button
			 */
			showLoading: function() {
				$("#repository-resources").html('<div class="text-center">' + wisply.getLoadingImage("medium") + '</div>');
			},
			/**
			 * Is is called when resources has came from the server. It shows them, activate the listeners and update the GUI
			 * @param  {string} html The HTML from server
			 */
			fired_NewResources: function(html) {
				this.changeResources(html);
				this.activateListeners();
				this.updateGUI();
			},
			/**
			 * It activates the listenrs for the next and previous buttons
			 */
			activateListeners: function() {
				var instance = this;
				/**
				 * It sets the listenes for next and previous buttons. In case they have a disabled attribute, it does not call their listeners
				 */
				function initButtons() {
					$(".next").click(function(event) {
						event.preventDefault();
						if (!$(this).hasClass("disabled")) {
							instance.showNext();
						}
					});
					$(".previous").click(function(event) {
						event.preventDefault();
						if (!$(this).hasClass("disabled")) {
							instance.showPrevious();
						}
					});
				}
				initButtons();
			},
			/**
			 * It shows the next resources
			 * @param  {event} event The event which has been generated when the button has been clicked
			 */
			showNext: function() {
				var newMin;
				newMin = parseInt(this.min, 10) + this.resourcePerPage;
				this.changeMin(newMin);
				this.goUp();
				this.updateHash();
			},
			/**
			 * It gets the previous resources
			 */
			showPrevious: function() {
				var newMin;
				newMin = parseInt(this.min, 10) - this.resourcePerPage;
				this.changeMin(newMin);
				this.goUp();
				this.updateHash();
			},
			/**
			 * It changes the min value (The value from which the resources are displayed)
			 * @param  {string} newValue The new value
			 */
			changeMin: function(newValue) {
				var value = 0;
				if ((!isInt(newValue)) || !newValue || newValue === "" || parseInt(newValue, 10) < 0) {
					value = 0;
				} else {
					value = newValue;
				}
				this.min = parseInt(value, 10);
			},
			/**
			 * It updates the hash of the page
			 */
			updateHash: function() {
				var instance = this;
				/**
				 * It returns the string for the "list" verb
				 * @return {string} The string for the "list" verb
				 */
				function getList() {
					return "list" + instance.delimitator.insideVerb + instance.min + instance.delimitator.parameter + instance.resourcePerPage;
				}
				/**
				 * It returns a string which holds all the verbs
				 * @return {string} All the verbs as a string
				 */
				function getVerbs() {
					var verbs = [];
					verbs.push(getList());
					return verbs.join(instance.delimitator.verb);
				}
				window.location.hash = getVerbs();
			},
			/**
			 * It is called when the hash of the page has been changed. It gets all the verbs from the hash and updates the page
			 */
			hashChanged: function() {
				var h = window.location.hash,
					hash = h.substr(h.indexOf('#') + 1),
					instance = this;
				/**
				 * It extracts the verbs from a string and stores them inside this.verbs variable
				 * @param  {string} URL The entire string
				 */
				function extractVerbs(URL) {
					var ret = [],
						verbs = URL.split(instance.delimitator.verb);
					var i, verb, extracted;
					for (i = 0; i < verbs.length; i++) {
						verb = verbs[i];
						extracted = extractElements(verb);
						ret.push(extracted);
					}
					return ret;
				}
				/**
				 * It extracts the elements of a verb and its name and returns them
				 * @param  {string} URLString The string which contains the verb
				 * @return {object} An object which holds the name of the verb and its paramters
				 */
				function extractElements(URLString) {
					var elements = {},
						el;
					el = URLString.split(instance.delimitator.insideVerb);
					elements.verb = el[0];
					if (el[1]) {
						elements.parameters = el[1].split(instance.delimitator.parameter);
					}
					return elements;
				}
				this.verbs = extractVerbs(hash);
				this.updateListVerb();
				this.getResources();
			},
			/**
			 * It returns a verb from the list of verbs
			 * @param  {string} name The name of the verb
			 */
			getVerb: function(name) {
				var verb,
					instance = this,
					verbs = instance.verbs;
				for (verb in verbs) {
					if (verbs.hasOwnProperty(verb)) {
						if (verbs[verb].verb === name) {
							return verbs[verb];
						}
					}
				}
			},
			/**
			 * It updates the "list" verb
			 */
			updateListVerb: function() {
				var verb = this.getVerb("list");
				if (verb) {
					this.changeMin(verb.parameters[0]);
					this.changeResourcesPerPage(verb.parameters[1]);
				}
			},
			/**
			 * It changes the number of resources displayed per page
			 * @param  {number} newValue The new value
			 */
			changeResourcesPerPage: function(newValue) {
				this.resourcePerPage = parseInt(newValue, 10);
			},
			/**
			 * It updates the next and previous buttons
			 */
			updateGUI: function() {
				var instance = this;
				/**
				 * If there are no more resources, it disables the previous button. Otherwise, it enables it
				 */
				function updatePreviousButton() {
					if (instance.min < instance.resourcePerPage) {
						$(".previous").addClass("disabled");
					} else {
						$(".previous").removeClass("disabled");
					}
				}
				/**
				 * If there are no more resources, it disables the next button. Otherwise, it enables it
				 */
				function updateNextButton() {
					if (instance.min + instance.resourcePerPage >= parseInt(instance.repository.totalRecords, 10)) {
						$(".next").addClass("disabled");
					} else {
						$(".next").removeClass("disabled");
					}
				}

				function updateButtons() {
					updatePreviousButton();
					updateNextButton();
				}
				updateButtons();
			},
			/**
			 * It takes the user up to the list of resources
			 */
			goUp: function() {
				var listPosition = parseInt($("#repository-before-resources").offset().top, 10) - 70;
				$('html, body').animate({
					scrollTop: listPosition,
				}, 100);
			},
			/**
			 * It changes the content of the DIV which holds the list of resources
			 * @param  {string} html The new HTML code for the DIV
			 */
			changeResources: function(html) {
				$("#repository-resources").html(html);
			},
			/**
			 * It changes the content of the UP DIV which holds the options
			 * @param  {string} html The new HTML code for the DIV
			 */
			updateTop: function(html) {
				$("#repository-top").html(html);
			},
			/**
			 * It changes the content of the BOTTOM DIV which holds the options
			 * @param  {string} html The new HTML code for the DIV
			 */
			updateBottom: function(html) {
				$("#repository-bottom").html(html);
			}
		};
	/**
	 * It checks if a string is int
	 * @param  {string}  value The value of the string
	 * @return {Boolean} True if the string is an int, false otherwise
	 */
	function isInt(value) {
		return !isNaN(value) && parseInt(Number(value)) == value && !isNaN(parseInt(value, 10));
	}
	return {
		Manager: Manager
	};
};
$(document).ready(function() {
	"use strict";
	var module = new PublicRepository();
	wisply.publicRepositoryModule = module;
	wisply.publicRepositoryModule.manager = new module.Manager(server.repository);
	wisply.publicRepositoryModule.manager.init();
});
