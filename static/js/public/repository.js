/* global $, wisply, server */
/**
 * @file Encapsulates the functionality for displaying the contents of a repository for the public
 * @author Cristian Sima
 */
/**
 * @namespace PublicRepositoryModule
 */
var PublicRepositoryModule = function() {
	'use strict';
	// credits http://stackoverflow.com/a/881147/2415167
	String.prototype.count = function(s1) {
		return (this.length - this.replace(new RegExp(s1, "g"), '').length) / s1.length;
	};
	// credits http://stackoverflow.com/a/646643/2415167
	if (typeof String.prototype.startsWith != 'function') {
		String.prototype.startsWith = function(str) {
			return this.indexOf(str) === 0;
		};
	}
	/**
	 * The constructor sets the default values
	 * @memberof PublicRepositoryModule
	 * @class Manager
	 * @classdesc It encapsulets the functionality for the public repository
	 * @param {object} currentRepository Represents the current repository
	 */
	var Manager = function Manager(currentRepository) {
		this.repository = currentRepository;
		this.min = 0;
		this.resourcesPerPage = 15;
		// They are used to extract the verbs and parameters from the hash
		this.delimitator = {
			verb: "*",
			parameter: "-",
			insideVerb: "|",
		};
		this.resourceFocused = -1;
		this.topGUI = new TopGUI(this);
		this.sideGUI = new SideGUI(this);
		this.bottomGUI = new BottomGUI(this);
	};
	Manager.prototype =
		/** @lends PublicRepositoryModule.Manager */
		{
			/**
			 * It activates the listeners and adds the shortcuts
			 */
			init: function() {
				var instance = this;
				/**
				 * It removes the old listener for change hash and sets one which calls hashChanged method
				 */
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
					}, {
						"type": "keyup",
						"key": "Down",
						"callback": function() {
							instance.focus(instance.resourceFocused + 1);
						},
						"description": function() {
							return "Focuses the first resource or focuses the next resource.";
						}(),
						"overwrites": true,
					}, {
						"type": "keyup",
						"key": "Up",
						"callback": function() {
							instance.focus(instance.resourceFocused - 1);
						},
						"description": function() {
							return "Focuses the previous resource.";
						}(),
						"overwrites": true,
					}];
					wisply.shortcutManager.activate(shortcuts);
				}

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
				initHash();
				initKeys();
				initButtons();
				this.initGUI();
				this.updateGUI();
				this.hashChanged();
			},
			/**
			 * It calls the GUIs' init method
			 */
			initGUI: function() {
				this.topGUI.init();
				this.sideGUI.init();
				this.bottomGUI.init();
			},
			/**
			 * It focuses the index resource
			 * @param  {number} index The index is from 0 to n-1 resourcesPerPage
			 */
			focus: function(index) {
				if (index + 1 >= this.resourcesPerPage) {
					index = this.resourcesPerPage - 1;
				}
				if (index > -1) {
					this.resourceFocused = index;
					$("#listOfRecords").find("a").eq(index).focus();
				}
			},
			/**
			 * It resets the focus item. Private method
			 */
			_resetFocus: function() {
				this.resourceFocused = -1;
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
					url: "/api/repository/resources/" + this.repository.id + "/get/" + this.min + "/" + this.resourcesPerPage,
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
			 * It shows the top div, the buttons and sets the wisply loading logo
			 */
			showLoading: function() {
				$("#repository-top, .next, .previous, #repository-bottom, #repository-side").hide();
				$("#repository-resources").html('<div class="text-center"><br /><br /><br />' + wisply.getLoadingImage("medium") + '</div>');
			},
			/**
			 * Is is called when resources has came from the server. It shows them, activate the listeners and update the GUI
			 * @param  {string} html The HTML from server
			 */
			fired_NewResources: function(html) {
				this.changeResources(html);
				this.updateGUI();
				this._resetFocus();
			},
			/**
			 * It shows the next resources
			 * @param  {event} event The event which has been generated when the button has been clicked
			 */
			showNext: function() {
				var instance = this;
				if (instance.min + instance.resourcesPerPage < parseInt(instance.getCurrentTotalNumber(), 10)) {
					var newMin;
					newMin = parseInt(this.min, 10) + this.resourcesPerPage;
					this.changeStart(newMin);
				}
			},
			/**
			 * It shows the last page for the repository
			 */
			showLastPage: function() {
				var total = this.getCurrentTotalNumber(),
					resourcesPerPage = this.resourcesPerPage,
					numberOfPages = Math.round(total / resourcesPerPage),
					newStart = (numberOfPages - 1) * resourcesPerPage;
				this.changeStart(newStart);
			},
			/**
			 * It gets the previous resources
			 */
			showPrevious: function() {
				var instance = this;
				if (instance.min >= instance.resourcesPerPage) {
					var newMin;
					newMin = parseInt(this.min, 10) - this.resourcesPerPage;
					this.changeStart(newMin);
				}
			},
			/**
			 * It changes the min value (The value from which the resources are displayed). It is a private method
			 * @param  {string} newValue The new value
			 */
			_changeStart: function(newValue) {
				var value = 0;
				if ((!isInt(newValue)) || !newValue || newValue === "" || parseInt(newValue, 10) < 0) {
					value = 0;
				} else {
					value = newValue;
				}
				this.min = parseInt(value, 10);
			},
			/**
			 * It changes the min value (The value from which the resources are displayed). It goes up and updates the hash.
			 * @param  {string} newValue The new value
			 */
			changeStart: function(newValue) {
				this._changeStart(newValue);
				this.goUp();
				this.updateHash();
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
					return "list" + instance.delimitator.insideVerb + instance.min + instance.delimitator.parameter + instance.resourcesPerPage;
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
			/**
			 * It removes the collection and sets the min to 0
			 */
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
				/**
				 * It gets the min and resourcesPerPage from the verb and passes them to the manager
				 */
				function updateList() {
					var verb = instance.getVerb("list");
					if (verb) {
						instance._changeStart(verb.parameters[0]);
						instance._changeResourcesPerPage(verb.parameters[1]);
					}
				}
				/**
				 * It changes the collection from the verb to the manager
				 */
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
			 * It can be used to change the number of resources displayed per page. It updates the hash
			 * @param  {number} newValue The new value
			 */
			changeResourcesPerPage: function(newValue) {
				this._changeResourcesPerPage(newValue);
				this.updateHash();
			},
			/**
			 * It changes the number of resources displayed per page. It is a private method
			 * @param  {number} newValue The new value
			 */
			_changeResourcesPerPage: function(newValue) {
				this.resourcesPerPage = parseInt(newValue, 10);
			},
			/**
			 * It updates the next and previous buttons
			 */
			updateGUI: function() {
				this.topGUI.update();
				this.sideGUI.update();
				this.bottomGUI.update();
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
	 * @memberof PublicRepositoryModule
	 * @class TopGUI
	 * @classdesc It encapsulets the functionality for top GUI
	 * @param {PublicRepositoryModule.Manager} manager Is a reference to the manager
	 */
	var TopGUI = function TopGUI(manager) {
		this.element = $("#repository-top");
		this.manager = manager;
	};
	TopGUI.prototype =
		/** @lends PublicRepositoryModule.TopGUI */
		{
			/**
			 * Does nothing. It is here for cosistency (the other GUI has an init method)
			 */
			init: function() {},
			/**
			 * It updates the Top GUI elements
			 */
			update: function() {
				var gui = this,
					manager = gui.manager;
				/**
				 * If there are no more resources, it disables the previous button. Otherwise, it enables it
				 */
				function updatePreviousButton() {
					if (manager.getCurrentTotalNumber() === 0 || manager.min < manager.resourcesPerPage) {
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
					if (manager.getCurrentTotalNumber() === 0 || (manager.min + manager.resourcesPerPage >= parseInt(manager.getCurrentTotalNumber(), 10))) {
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
					$("#repository-top, #repository-bottom, #repository-side").show();
				}
				/**
				 * It updates the description of the top DIV.
				 */
				function updateTop() {
					/**
					 * It returns the description for the list of repositories
					 * @return {string} The description for the list of repositories
					 */
					function getListDescription() {
						/**
						 * It returns the description of the list when there are repositories
						 * @return {string} The HTML description of the list
						 */
						function getListText() {
							/**
							 * It shows when the start is zero
							 * @return {string} The description when start is zero
							 */
							function getStartZero() {
								function describeCollection() {
									var text, total = manager.getCurrentTotalNumber();
									if (total < manager.resourcesPerPage) {
										if (total === 1) {
											text = "This collection contains only one resource";
										} else {
											text = "Showing all " + total + " resources";
										}
									} else {
										text = "Showing first " + manager.resourcesPerPage + " resources of a total number of " + manager.getCurrentTotalNumber();
									}
									return text;
								}
								var text = "";
								if (manager.collection) {
									text = describeCollection();
								} else {
									text = "Last resources on " + manager.repository.name + ":";
								}
								return text;
							}
							/**
							 * It shows when the start is not zero
							 * @return {string} The description when start is not zero
							 */
							function getStartNotZero() {
								var text = "";
								if (start + manager.resourcesPerPage >= manager.getCurrentTotalNumber()) {
									difference = manager.getCurrentTotalNumber() - start;
									text = "Showing last " + difference + " resources of total " + manager.getCurrentTotalNumber() + "";
								} else {
									text = "Showing " + manager.resourcesPerPage + " resources from " + start + " to " + end + " of a total number of " + manager.getCurrentTotalNumber() + "";
								}
								return text;
							}
							var text = "",
								start = manager.min,
								difference,
								end = manager.min + manager.resourcesPerPage;
							if ((start === 0)) {
								text += getStartZero();
							} else {
								text += getStartNotZero();
							}
							return text;
						}
						var listText = "",
							html = "";
						if (manager.getCurrentTotalNumber() === 0) {
							listText += "There are no resources.";
						} else {
							listText += getListText();
						}
						html += "<span class='text-muted'>";
						html += listText;
						html += "</span>";
						return html;
					}
					/**
					 * It returns the description for collections
					 * @return {string} The description for collections
					 */
					function getCollection() {
						/**
						 * It transforsm the parents of the current category into labels
						 * @return {string} HTML code containing the labels for the category's parents
						 */
						function getLabels() {
							/**
							 * It returns the HTML label for a given collection. In case the collection's name is longer than 20 characters, it cuts it to 20, appends "..." and shows a tooltip
							 * @param  {object} collection The collection which will be transformed into a tooltip
							 * @return {string} The HTML label for the collection
							 */
							function getLabel(collection) {
								var text = "",
									tooltip = "",
									name = collection.Path,
									maxCharacters = 20;
								if (elements.length != 1) {
									if (name.length > maxCharacters) {
										name = name.substring(0, maxCharacters) + "...";
										tooltip = 'data-toggle="tooltip" data-original-title="' + collection.Path + '"';
									}
								}
								text = '<span ' + tooltip + ' data-id="' + collection.ID + '" class="hover label label-info set-collection">' + name + '</span> ';
								return text;
							}
							/**
							 * It returns the HTML for remove collection button
							 */
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
									labelText += "<span class='text-muted glyphicon glyphicon-menu-right'></span> ";
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
					html += "<div id='collection'>" + getCollection() + "</div>";
					html += "<div>" + getListDescription() + "</div>";
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
	 * @memberof PublicRepositoryModule
	 * @class SideGUI
	 * @classdesc It encapsulets the functionality for side GUI
	 * @param {PublicRepositoryModule.Manager} manager Is a reference to the manager
	 */
	var SideGUI = function SideGUI(manager) {
		this.element = $("#repository-side");
		this.manager = manager;
		this.showAll = false;
		this.hideEmptyCollections = true;
	};
	SideGUI.prototype =
		/** @lends PublicRepositoryModule.SideGUI */
		{
			/**
			 * Does thing. It is present for consistency
			 */
			init: function() {},
			/**
			 * It updates the GUI
			 */
			update: function() {
				var instance = this,
					div = "";
				/**
				 * Returns the HTML code for the current collections
				 * @return {string} HTML code
				 */
				function getCollectionsHTML() {
					/**
					 * It returns the collections to be processed. In case there is a current collection, it retunrs the first children of it. Otherwise, if there is no collection selected, it returns all of them. Also, it trims the empty ones (if the hideEmptyCollections is true)
					 * @return {array} Collections to be processed
					 */
					function getCollectionsToProcess() {
						/**
						 * Returns the level of a collection. The level is the number of ":" plus one inside the Spec
						 * @param  {object} collection The collection
						 * @return {string} The level of collection
						 */
						function getLevel(collection) {
							return parseInt(collection.Spec.count(":"), 10) + 1;
						}
						/**
						 * It checks if the child is a direct children of a parent collection. It is a direct children if the child level plus one equals the parent and the child's Spec starts with parent's one
						 * @param  {object}  parent The parent collection
						 * @param  {object}  child  The child collection
						 * @return {Boolean} True fi the child is a direct children of the parent
						 */
						function isDirectChildrenOf(parent, child) {
							return child.Spec.startsWith(parent.Spec) && (getLevel(parent) + 1) === getLevel(child);
						}
						/**
						 * It returns a list of collections which are not empty. An empty collection has no resources
						 * @param  {array} collections The collection to be processed
						 * @return {array} The collections array, without the empty collections
						 */
						function removeEmpty(collections) {
							var i, toReturn = [];
							for (i = 0; i < collections.length; i++) {
								if (parseInt(collections[i].NumberOfResources, 10) !== 0) {
									toReturn.push(collections[i]);
								}
							}
							return toReturn;
						}
						/**
						 * It returns the array of collections which are direct children of a parent collection
						 * @param  {array} collections       All the collections
						 * @param  {object} currentCollection The parent collection
						 * @return {array}
						 */
						function getNextLevel(collections, currentCollection) {
							var toProcess = [],
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
						/**
						 * It sorts an array of collection by Name ASC
						 * @param  {collections} collections The array of collections to be sorted
						 * @return {array} The sorted array
						 */
						function sortCollections(collections) {
							var algorithm = function(a, b) {
								if (a.Name < b.Name) {
									return -1;
								}
								if (a.Name > b.Name) {
									return 1;
								}
								return 0;
							};
							return collections.sort(algorithm);
						}
						function decideWhichCollections() {
							var allCollections = instance.manager.repository.collections,
								toProcess,
								currentCollection = instance.manager.collection;
							if (!currentCollection || instance.showAll) {
								toProcess = allCollections;
							} else {
								toProcess = getNextLevel(allCollections, currentCollection);
							}
							return toProcess;
						}
						var	toProcess;

						toProcess = decideWhichCollections();

						// remove empty
						if (instance.hideEmptyCollections) {
							toProcess = removeEmpty(toProcess);
						}
						return sortCollections(toProcess);
					}
					/**
					 * It retunrs the HTML code for am array of collections
					 * @param  {array} collections The array of collections
					 * @return {string} The HTML code for the array
					 */
					function getCollections(collections) {
						/**
						 * It returns the HTML code for the collection
						 * @param  {object} collection The collection to be processed
						 * @return {string}
						 */
						function getCollection(collection) {
							var collectionHTML = "";
							/**
							 * It returns the HTML description of the collection
							 * @param  {string} description The value of description
							 * @return {string} The HTML description of the collection
							 */
							function getDescription(description) {
								return '<p class="list-group-item-text">' + description + '</p>';
							}
							/**
							 * It returns the div which holds the number of resources
							 * @param  {object} collection The collection
							 * @return {string} The HTML code for the DIV which holds the number of resources
							 */
							function getRightDiv(collection) {
								return '<div class="col-lg-3 col-md-3 col-sm-3 col-xs-3 text-right"><span class="text-right badge"> ' + collection.NumberOfResources + "</span></div>";
							}
							/**
							 * It returns the div which holds the name of collection
							 * @param  {object} collection The collection
							 * @return {string} The HTML code for the DIV which holds the name of collection
							 */
							function getLeftDiv(collection) {
								return '<div class="col-lg-9 col-md-9 col-sm-9 col-xs-9"><span class="h6">' + collection.Path + '</span></div>';
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
					var collectionsToProcess;
					collectionsToProcess = getCollectionsToProcess();
					return getCollections(collectionsToProcess);
				}
				/**
				 * It returns the HTML for the top DIV. This DIV may contain options regarding the collections
				 * @return {string} The HTML code for the top DIV
				 */
				function getTopDiv() {
					/**
					 * It returns the HTML code for the "show all" button
					 * @return {string} HTML code for the "show all"
					 */
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
					/**
					 * It returns the HTML code for the "Disply empty" button
					 * @return {string} The HTML code for the "Disply empty"
					 */
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
				/**
				 * It returns the HTML code for the entire sideGUI
				 * @return {string} HTML code for the entire sideGUI
				 */
				function getHTML() {
					var div = "";
					div += getTopDiv();
					div += getCollectionsHTML();
					return div;
				}
				/**
				 * It activates all the listeners for the sideGUI
				 */
				function activate() {
					/**
					 * It activates the listeners for "Show all" and "Hide empty" buttons
					 */
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
			/**
			 * It shows/hides all the collections
			 */
			toggleAllCollections: function() {
				this.showAll = !this.showAll;
				this.update();
			},
			/**
			 * It shows/hides the empty collections
			 */
			toggleEmptyCollections: function() {
				this.hideEmptyCollections = !this.hideEmptyCollections;
				this.update();
			}
		};
	/**
	 * Contains the functionality for bottom GUI
	 * @memberof PublicRepositoryModule
	 * @class BottomGUI
	 * @classdesc It encapsulets the functionality for bottom GUI
	 * @param {PublicRepositoryModule.Manager} manager Is a reference to the manager
	 */
	var BottomGUI = function BottomGUI(manager) {
		this.element = $("#repository-bottom");
		this.manager = manager;
		this.showMoreOptions = false;
	};
	BottomGUI.prototype =
		/** @lends PublicRepositoryModule.BottomGUI */
		{
			/**
			 * It activates the listeners for the bottom GUI
			 */
			init: function() {
				var instance = this;
				/**
				 * It sets a listener for the more option button
				 */
				function initShowMore() {
					$(".show-more-options").click(function() {
						instance.toggleShowMoreOptions();
					});
				}
				/**
				 * It sets a listener for the resource per page select
				 */
				function initResourcesPerPage() {
					$(".change-resources-per-page").on('change', function() {
						instance.manager.changeResourcesPerPage(this.value);
					});
				}
				/**
				 * It activates the listeners for last and first buttons
				 */
				function initButtons() {
					$(".show-first-resources").click(function(event) {
						event.preventDefault();
						instance.manager.changeStart(0);
					});
					$(".show-last-resources").click(function(event) {
						event.preventDefault();
						instance.manager.showLastPage();
					});
				}
				initShowMore();
				initResourcesPerPage();
				initButtons();
			},
			/**
			 * It updates the bottom GUI
			 */
			update: function() {
				var instance = this;
				/**
				 * It returns the HTML code for the "show more options" button
				 * @return {string} The HTML code for the "show more options" button
				 */
				function getMoreOptions() {
					var advanceHTML = "",
						text = "";
					if (instance.showMoreOptions) {
						text = "<span class='glyphicon glyphicon-remove'></span> Hide options";
					} else {
						text = "Show more options";
					}
					advanceHTML = "<span class='show-more-options hover text-muted'>" + text + "</span>";
					return advanceHTML;
				}
				/**
				 * It returns the HTML for the options
				 * @return {string} The HTML for the options
				 */
				function getOptions() {
					/**
					 * It returns the HTML for resources per page option
					 * @return {string} The HTML for resources per page option
					 */
					function getResourcesPerPage() {
						/**
						 * It returns the HTML for the select object
						 * @param  {number} currentNumber The current number of resources
						 * @return {string} The HTML for the select object
						 */
						function getSelectHTML(currentNumber) {
							var selectHTML = "",
								options = [5, 15, 25, 50, 100],
								i, number, selected;
							selectHTML += "<select class='change-resources-per-page'>";
							for (i = 0; i < options.length; i++) {
								number = options[i];
								if (parseInt(number, 10) === parseInt(currentNumber, 10)) {
									selected = "selected";
								} else {
									selected = "";
								}
								selectHTML += "<option " + selected + ">" + number + "</option>";
							}
							selectHTML += "</select>";
							return selectHTML;
						}
						var perPageHTML = "",
							number = instance.manager.resourcesPerPage;
						perPageHTML += "Display " + getSelectHTML(number) + " resources per page.";
						return perPageHTML;
					}
					/**
					 * Returns the HTMl code for the buttons
					 * @return {string} The HTML code for the buttons
					 */
					function getButtons() {
						/**
						 * It returns the HTML code for the first page button
						 * @return {string} The HTML code for the first page button
						 */
						function getFirst() {
							var text = "",
								total = instance.manager.getCurrentTotalNumber(),
								start = instance.manager.min,
								resourcesPerPage = instance.manager.resourcesPerPage;
							if (total < resourcesPerPage || (start === 0)) {
								text = "";
							} else {
								text = '<li class="previous"><a href="#" class="show-first-resources">← First page</a></li>';
							}
							return text;
						}
						/**
						 * It returns the code for the last page button
						 * @return {string} The HTML code for the last page button
						 */
						function getLast() {
							var text = "",
								total = instance.manager.getCurrentTotalNumber(),
								start = instance.manager.min,
								resourcesPerPage = instance.manager.resourcesPerPage;
							if (start + resourcesPerPage >= total) {
								text = "";
							} else {
								text = '<li class="next"><a href="#" class="show-last-resources">Last page →</a></li>';
							}
							return text;
						}
						var buttonsHTML = "";
						buttonsHTML += "<ul class='pager'>";
						buttonsHTML += getFirst();
						buttonsHTML += getLast();
						buttonsHTML += "</ul>";
						return buttonsHTML;
					}
					var text = "",
						visibility;
					visibility = (!instance.showMoreOptions ? "style='display:none'" : "");
					text += "<div " + visibility + " class='well' >";
					text += getResourcesPerPage();
					text += "<hr />";
					text += getButtons();
					text += "</div>";
					return text;
				}
				if(instance.manager.getCurrentTotalNumber() !== 0) {
					var bottomGUIHTML = "<hr />";
					bottomGUIHTML += getMoreOptions();
					bottomGUIHTML += getOptions();
					this.element.html(bottomGUIHTML);
					this.init();
				} else {
					this.element.html("");
				}
			},
			/**
			 * It shows/hides the div which contains more options
			 */
			toggleShowMoreOptions: function() {
				this.showMoreOptions = !this.showMoreOptions;
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
	var module = new PublicRepositoryModule();
	wisply.loadModule("public-repository", module);
});
