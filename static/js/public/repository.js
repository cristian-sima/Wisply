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
					// distroy the main hashchange
					$(window).unbind("hashchange");
					// add a listener for this manager
					$(window).on('hashchange', function() {
						instance.hashChanged();
					});
				}
				/**
				 * It adds the keys' shortcuts
				 */
				function initKeys() {
					var instance,
						shortcuts = [{
						"type": "keyup",
						"key": "Ctrl+left",
						"callback": function() {
							wisply.publicRepositoryModule.manager.showPrevious();
						},
						"description": "Receives the next page of resources",
					}, {
						"type": "keyup",
						"key": "Ctrl+right",
						"callback": function() {
							wisply.publicRepositoryModule.manager.showNext();
						},
						"description": "Receives the previous page of resources",
					}];
					wisply.shortcutManager.activate(shortcuts);
				}
				/**
				 * It sets a listener for each collection button
				 */
				function initCollections() {
					$(".set-collection").click(function(event) {
							event.preventDefault();
							var id = $(this).data("id");
							instance.setCollection(id);
					});
				}
				initHash();
				initKeys();
				initCollections();
				this.updateGUI();
        this.hashChanged();
			},
			/**
			 * It loads the resources from server according to the current settings
			 */
			getResources: function() {
				var instance = this,
					collection = getCollection();
				/**
				 * Returns the ID of the current collection or an empty string if there are no collections
				 * @return {number} The ID of current collection
				 */
				function getCollection() {
					if(instance.collection) {
						return instance.collection.Spec;
					}
					return "";
				}
				this.showLoading();
				$.ajax({
					url: "/api/repository/resources/" + this.repository.id + "/get/" + this.min + "/" + this.resourcePerPage,
					data: {
						"collection": getCollection(),
						"format": "html",
					},
					success: function(html) {
						instance.fired_NewResources(html);
					}
				});
			},
			/**
			 * It shows the wisply loading button
			 */
			showLoading: function() {
				$(".next, .previous, #repository-top").hide();
				$("#repository-resources").html('<div class="text-center"><br /><br /><br />' + wisply.getLoadingImage("medium") + '</div>');
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
				var instance = this;
				if (instance.min + instance.resourcePerPage < parseInt(instance.getCurrentTotalNumber(), 10)) {
					var newMin;
					newMin = parseInt(this.min, 10) + this.resourcePerPage;
					this.changeMin(newMin);
					this.goUp();
					this.updateHash();
				}
			},
			/**
			 * It gets the previous resources
			 */
			showPrevious: function() {
				var instance = this;
				if (instance.min >= instance.resourcePerPage) {
					var newMin;
					newMin = parseInt(this.min, 10) - this.resourcePerPage;
					this.changeMin(newMin);
					this.goUp();
					this.updateHash();
				}
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
				 * It returns the description of the verb collection
				 * @return {string} The description of the verb collection
				 */
				function getCollection() {
					if(instance.collection) {
						return "collection" + instance.delimitator.insideVerb + instance.collection.ID;
					}
					return "";
				}
				/**
				 * It returns a string which holds all the verbs
				 * @return {string} All the verbs as a string
				 */
				function getVerbs() {
					var verbs = [];
					verbs.push(getList());
					verbs.push(getCollection());
					return verbs.join(instance.delimitator.verb);
				}
				window.location.hash = getVerbs();
			},
			/**
			 * The current total number is the total number of the resources inside a collection, or the total number of resources from the repository in case there is no repository selected
			 * @return {number} The current total number
			 */
			getCurrentTotalNumber: function() {
				if(this.collection) {
					return this.collection.NumberOfResources;
				}
				return this.repository.totalResources;
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
				this.updateVerbs();
				this.getResources();
			},
			/**
			 * It gets the object of the collection, it sets it, goes up and updates the hash
			 * @param  {number} id The ID of the collection
			 */
			setCollection: function(id) {
					var collection = this.repository.getCollection(id);
					this.collection = collection;
					this.goUp();
					this.updateHash();
			},
			removeCollection: function() {
				var instance = this;
				instance.min = 0;
				instance.collection = undefined;
				instance.updateHash();
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
			updateVerbs: function() {
				var instance = this;
				function updateList() {
						var verb = instance.getVerb("list");
					if (verb) {
						instance.changeMin(verb.parameters[0]);
						instance.changeResourcesPerPage(verb.parameters[1]);
					}
				}
				function updateCollection() {
					var verb = instance.getVerb("collection");
					if (verb) {
						var collection = instance.repository.getCollection(verb.parameters[0]);
						instance.collection = collection;
					}
				}
				updateList();
				updateCollection();
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
					if (instance.min + instance.resourcePerPage >= parseInt(instance.getCurrentTotalNumber(), 10)) {
						$(".next").addClass("disabled");
					} else {
						$(".next").removeClass("disabled");
					}
				}
				/**
				 * It updates the next and previous button
				 */
				function updateButtons() {
					updatePreviousButton();
					updateNextButton();
				}
				/**
				 * It shows the buttons
				 */
				function showElements() {
						$(".next, .previous, #repository-top").show();
				}
				/**
				 * It updates the description of the top DIV.
				 */
				function updateTop() {
					/**
					 * It returns the description for the list of repositories
					 * @return {string} The description for the list of repositories
					 */
					function getShowing() {
						var start = instance.min, difference,
							end = instance.min + instance.resourcePerPage,
							text = "", html = "";
						if ((start === 0)) {
							if (instance.collection) {
								text = "Showing first " + instance.resourcePerPage + " resources of a total number of " + instance.getCurrentTotalNumber();
							} else {
								text = "Last resources:";
							}
						} if(start + instance.resourcePerPage >= instance.getCurrentTotalNumber() ) {
							difference = instance.getCurrentTotalNumber() - start;
							text = "Showing last " +  difference + " resources of total " + instance.getCurrentTotalNumber() + "";
						} else {
							text = "Showing " +  instance.resourcePerPage + " resources from " + start + " to " + end + " of " + instance.getCurrentTotalNumber() + "";
						}
						 html += "<span class='text-muted'>";
						 html += text;
						 html += "</span>";
						 return html;
					}
					/**
					 * It returns the description for collections
					 * @return {string} The description for collections
					 */
					function getCollection() {
						var text = "";
						if (instance.collection) {
							text = '<span class="label label-info">' + instance.collection.Name + '</span> <a data-toggle="tooltip" data-id="' + instance.collection.id + '" id="remove-collection" href="#" data-original-title="Remove collection"><span class="text-danger glyphicon glyphicon-remove"></span></a>';
						} else {
							text = "<br />";
						}
						return text;
					}
					var html = "";
					html += "<div>";
					html += "<div>" + getShowing() + "</div>";
					html += "<div id='collection'>" + getCollection() + "</div>";
					html += "<br /></div>";
					$("#repository-top").html(html);
				}
				/**
				 * It sets a listener for the remove collection button
				 */
				function initCollection() {
					$("#remove-collection").click(function(event) {
						event.preventDefault();
						instance.removeCollection();
					});
				}
				updateTop();
				updateButtons();
				showElements();
				wisply.activateTooltip();
				initCollection();
			},
			/**
			 * It takes the user up to the list of resources
			 */
			goUp: function() {
				var listPosition = parseInt($("#repository-before-resources").offset().top, 10) - 160;
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
