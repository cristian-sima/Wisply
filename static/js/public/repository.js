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
	String.prototype.count = function(s1) {
		return (this.length - this.replace(new RegExp(s1, "g"), '').length) / s1.length;
	};
	if (typeof String.prototype.startsWith != 'function') {
		// see below for better implementation!
		String.prototype.startsWith = function(str) {
			return this.indexOf(str) === 0;
		};
	}
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
		this.topGUI = new TopGUI(this);
		this.sideGUI = new SideGUI(this);
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
					var shortcuts = [{
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
				initHash();
				initKeys();
				this.initGUI();
				this.updateGUI();
				this.hashChanged();
			},
			initGUI: function() {
				this.topGUI.init();
				this.sideGUI.init();
			},
			/**
			 * It loads the resources from server according to the current settings
			 */
			getResources: function() {
				var instance = this;
				/**
				 * Returns the ID of the current collection or an empty string if there are no collections
				 * @return {number} The ID of current collection
				 */
				function getCollection() {
					if (instance.collection) {
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
				$("#repository-top, .next, .previous").hide();
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
					if (instance.collection) {
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
				if (this.collection) {
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
				this.min = 0;
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
				this.topGUI.update();
				this.sideGUI.update();
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
	 * Contains the functionality for top GUI
	 * @memberof PublicRepository
	 * @class TopGUI
	 * @classdesc It encapsulets the functionality for top GUI
	 * @param {PublicRepository.Manager} manager Is a reference to the manager
	 */
	var TopGUI = function TopGUI(manager) {
		this.element = $("#repository-top");
		this.manager = manager;
	};
	TopGUI.prototype =
		/** @lends PublicRepository.TopGUI */
		{
			init: function() {},
			update: function() {
				var gui = this,
					manager = gui.manager;
				/**
				 * If there are no more resources, it disables the previous button. Otherwise, it enables it
				 */
				function updatePreviousButton() {
					if (manager.getCurrentTotalNumber() === 0 || manager.min < manager.resourcePerPage) {
						$(".previous").addClass("disabled");
						$(".previous").hide();
					} else {
						$(".previous").removeClass("disabled");
						$(".previous").show();
					}
				}
				/**
				 * If there are no more resources, it disables the next button. Otherwise, it enables it
				 */
				function updateNextButton() {
					if (manager.getCurrentTotalNumber() === 0 || (manager.min + manager.resourcePerPage >= parseInt(manager.getCurrentTotalNumber(), 10))) {
						$(".next").addClass("disabled");
						$(".next").hide();
					} else {
						$(".next").removeClass("disabled");
						$(".next").show();
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
					$("#repository-top").show();
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
						var start = manager.min,
							difference,
							end = manager.min + manager.resourcePerPage,
							text = "",
							html = "";
						if (manager.getCurrentTotalNumber() === 0) {
							text = "There are no resources.";
						} else {
							if ((start === 0)) {
								if (manager.collection) {
									text = "Showing first " + manager.resourcePerPage + " resources of a total number of " + manager.getCurrentTotalNumber();
								} else {
									text = "Last resources:";
								}
							}
							if (start + manager.resourcePerPage >= manager.getCurrentTotalNumber()) {
								difference = manager.getCurrentTotalNumber() - start;
								if (difference === manager.getCurrentTotalNumber()) {
									if(difference === 1) {
										text = "There is only one resource";
									} else {
										text = "There are only " + difference + " resources";
									}
								} else {
									text = "Showing last " + difference + " resources of total " + manager.getCurrentTotalNumber() + "";
								}
							} else {
								text = "Showing " + manager.resourcePerPage + " resources from " + start + " to " + end + " of " + manager.getCurrentTotalNumber() + "";
							}
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
						function getLabels() {
							function getLabel(collection) {
								var text = "", tooltip="", name = collection.Name;
								if(elements.length != 1) {
									if (name.length > 20) {
										name = name.substring(0, 20) + "...";
										tooltip = 'data-toggle="tooltip" data-original-title="' + collection.Name + '"';
									}
								}
								text = '<span ' + tooltip + ' data-id="' + collection.ID + '" class="hover label label-info set-collection">' + name + '</span> ';
								return text;
							}

							function getLast() {
								var text = "";
								text = '<a data-toggle="tooltip" id="remove-collection" href="#" data-original-title="Remove collection"><span class="text-danger glyphicon glyphicon-remove"></span></a>';
								return text;
							}
							var labelText = "",
								currentSpec, parentSpec = "",
								i, parent, elements, currentSet = manager.collection.Spec,
								set;
							elements = currentSet.split(":");
							for (i = 0; i < elements.length; i++) {
								set = elements[i];
								currentSpec = parentSpec + set;
								parent = manager.repository.getBySpec(currentSpec);
								labelText += getLabel(parent);
								parentSpec = currentSpec + ":";
								if (i < elements.length - 1) {
									labelText += " <span class='text-muted glyphicon glyphicon-menu-right'></span> ";
								}
							}
							labelText += getLast();
							return labelText;
						}
						var text = "";
						if (manager.collection) {
							text += getLabels();
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
					manager.sideGUI.activateCollection();
				}
				/**
				 * It sets a listener for the remove collection button
				 */
				function initCollection() {
					$("#remove-collection").click(function(event) {
						event.preventDefault();
						manager.removeCollection();
					});
				}
				updateTop();
				updateButtons();
				showElements();
				wisply.activateTooltip();
				initCollection();
			}
		};
	/**
	 * Contains the functionality for side GUI
	 * @memberof PublicRepository
	 * @class SideGUI
	 * @classdesc It encapsulets the functionality for side GUI
	 * @param {PublicRepository.Manager} manager Is a reference to the manager
	 */
	var SideGUI = function SideGUI(manager) {
		this.element = $("#repository-side");
		this.manager = manager;
		this.showAll = false;
		this.hideEmptyCollections = true;
	};
	SideGUI.prototype =
		/** @lends PublicRepository.SideGUI */
		{
			init: function() {},
			update: function() {
				var instance = this,
					div = "";
				/**
				 * Returns the HTML code for the current collections
				 * @return {string} HTML code
				 */
				function getCollectionsHTML() {
					var collectionsToProcess;

					function getCollectionsToProcess() {
						function getLevel(collection) {
							return parseInt(collection.Spec.count(":"), 10) + 1;
						}

						function isDirectChildrenOf(parent, child) {
							return child.Spec.startsWith(parent.Spec) && (getLevel(parent) + 1) === getLevel(child);
						}

						function removeEmpty(collections) {
							var i, toReturn = [];
							for (i = 0; i < collections.length; i++) {
								if (parseInt(collections[i].NumberOfResources, 10) !== 0) {
									toReturn.push(collections[i]);
								}
							}
							return toReturn;
						}

						function getNextLevel(collections, currentCollection) {
							var toProcess = [],
								level = getLevel(currentCollection),
								i,
								checkCollection;
							for (i = 0; i < collections.length; i++) {
								checkCollection = collections[i];
								if (isDirectChildrenOf(currentCollection, checkCollection)) {
									toProcess.push(checkCollection);
								}
							}
							return toProcess;
						}
						var allCollections = instance.manager.repository.collections,
							toProcess,
							currentCollection = instance.manager.collection;
						if (!currentCollection || instance.showAll) {
							toProcess = allCollections;
						} else {
							toProcess = getNextLevel(allCollections, currentCollection);
						}
						// remove empty
						if (instance.hideEmptyCollections) {
							toProcess = removeEmpty(toProcess);
						}
						return toProcess.sort(function(a, b) {
							if (a.Name < b.Name) {
								return -1;
							}
							if (a.Name > b.Name) {
								return 1;
							}
							return 0;
						});
					}

					function getCollections(collections) {
						function getCollection(collection) {
							var collectionHTML = "";

							function getDescription(description) {
								return '<p class="list-group-item-text">' + description + '</p>';
							}

							function getRightDiv(collection) {
								return '<div class="col-lg-3 col-md-3 col-sm-3 col-xs-3 text-right"><span class="text-right badge"> ' + collection.NumberOfResources + "</span></div>";
							}

							function getLeftDiv(collection) {
								return '<div class="col-lg-9 col-md-9 col-sm-9 col-xs-9"><span class="h6">' + collection.Name + '</span></div>';
							}
							collectionHTML += '<a data-id="' + collection.ID + '" class="hover list-group-item set-collection">';
							collectionHTML += '<div class="row list-group-item-heading">';
							collectionHTML += getLeftDiv(collection);
							collectionHTML += getRightDiv(collection);
							collectionHTML += getDescription(collection.Description);
							collectionHTML += '</div>';
							collectionHTML += '</a>';
							return collectionHTML;
						}
						var html = "",
							i;
						for (i = 0; i < collections.length; i++) {
							html += getCollection(collections[i]);
						}
						if (html === "") {
							html = "<div class='text-center text-muted'>No collection available.</div>";
						}
						return html;
					}
					collectionsToProcess = getCollectionsToProcess();
					return getCollections(collectionsToProcess);
				}

				function getTopDiv() {
					function getLeft() {
						var htmlLeft = "";
						if (instance.manager.collection) {
							var text = "";
							if (!instance.showAll) {
								text = "Show all";
							} else {
								text = "Hide all";
							}
							htmlLeft = "<span href='' class='hover link show-all-collections text-muted'>" + text + "</span>";
						}
						return htmlLeft;
					}

					function getRight() {
						var text = "",
							htmlRight = "";
						if (instance.hideEmptyCollections) {
							text = "Display empty";
						} else {
							text = "Hide empty";
						}
						htmlRight = "<span href='' class='hover link hide-empty-collections text-muted'>" + text + "</span>";
						return htmlRight;
					}
					var div = "";
					div += "<div class='row'>";
					div += "<div class='col-lg-6 col-md-6 col-sm-6 col-xs-6 text-left'>" + getLeft() + "</div>";
					div += "<div class='col-lg-6 col-md-6 col-sm-6 col-xs-6 text-right'>" + getRight() + "</div>";
					div += "</div>";
					return div;
				}

				function getHTML() {
					var div = "";
					div += getTopDiv();
					div += getCollectionsHTML();
					return div;
				}

				function activate() {
					function activateTop() {
						$(".show-all-collections").click(function(event) {
							event.preventDefault();
							instance.toggleAllCollections();
						});
						$(".hide-empty-collections").click(function(event) {
							event.preventDefault();
							instance.toggleEmptyCollections();
						});
					}
					activateTop();
					instance.activateCollection();
				}
				div = getHTML();
				this.element.html(div);
				activate();
			},
			/**
			 * It sets a listener for each collection button
			 */
			activateCollection: function() {
				var instance = this;
				$(".set-collection").click(function(event) {
					event.preventDefault();
					var id = $(this).data("id");
					instance.manager.setCollection(id);
				});
			},
			toggleAllCollections: function() {
				this.showAll = !this.showAll;
				this.update();
			},
			toggleEmptyCollections: function() {
				this.hideEmptyCollections = !this.hideEmptyCollections;
				this.update();
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
