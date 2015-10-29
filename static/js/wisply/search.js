/* global $, wisply, Handlebars, Bloodhound */
var searchModule = {};
/**
 * @file Encapsulates the functionality for a search
 * @author Cristian Sima
 */
/**
 * Encapsulates the functionality for search
 * @namespace SearchModule
 */
var SearchModule = function() {
	'use strict';

	/**
	 * In case the length of the string is greater than the numberOfAllowedCharacters
	 * it cuts the string and appends "..."
	 * @param  {string} theString                 The string to be modified
	 * @param  {number} numberOfAllowedCharacters The number of characters to display
	 * @return {string} The modified string
	 */
	function cutString(theString, numberOfAllowedCharacters) {
		if (theString.length > numberOfAllowedCharacters) {
			var shortString = theString.substr(0, numberOfAllowedCharacters);
			theString = shortString + "...";
		}
		return theString;
	}
	// register a new function for the JavaScript template library
	// it is used to cut a string which is too long
	Handlebars.registerHelper('cutString', function(theString, numberOfAllowedCharacters) {
		theString = cutString(theString, numberOfAllowedCharacters);
		return new Handlebars.SafeString(theString);
	});
	/**
	 * It creates the search engine and activates the listeners
	 * @memberof SearchModule
	 * @class Field
	 * @param [object] o Contains the settings for the search field. The settings are: <br />
	 * selector  - the HTML selector for the element
	 * URL - The URL address for searches. The request will be o.URL/{{query}
	 * saveSearches - A boolean which indicates if the searches are stored in the cookies
	 * @classdesc Represents a search fields which implements the typeahead input
	 * @see https://twitter.github.io/typeahead.js/
	 */
	var Field = function Field(o) {
		this.selector = o.selector;
		this.URL = o.URL;
		this.saveSearches = o.saveSearches;
		this.cookies = new Cookies(this);
		this._init();
	};
	Field.prototype =
		/** @lends SearchModule.Field */
		{
			/**
			 * Creates the engines, loads the object and activates the listeners.
			 * It is a private method, called by constructor
			 */
			_init: function() {
				var instance = this;
				/**
				 * It creates the engine which manages the searches
				 * @return {object} The engine
				 */
				function createObject() {
					var settings, suggestionsEngine, searchesEngine;
					/**
					 * It activates the listeners for the typeahead engine
					 */
					function activateListeners() {
						// I used `event`, even if I do not use it in order to keep
						// the code cosistent
						/* jshint unused:false */
						instance.object.bind('typeahead:select', function(event, suggestion) {
							var that = this,
								instanceCopy = instance,
								suggestionCopy = suggestion;
							/**
							 * It is fired when a result has been selected and it contains a
							 * category. It check if the account is connected, if so it saves
							 * the query on the server.
							 */
							function resultSelected() {
								/**
								 * It saves a query on the server and then executes its action
								 * @param  {object} suggestion The query/suggestion object
								 */
								function saveQuery(suggestion) {
									var requestObject = {
										url: "/api/search/save/" + suggestion.Title,
										success: function() {
											window.location = suggestion.URL;
										},
									};
									wisply.executePostAjax(requestObject);
								}
								if (instanceCopy.saveSearches) {
									instanceCopy.cookies.saveSearchQuery(suggestion.Title);
								}
								if (wisply.connection) {
									saveQuery(suggestion);
								} else {
									window.location = suggestion.URL;
								}
							}
							if (suggestion.Category) {
								resultSelected();
							} else {
								$(this).typeahead('close');
								setTimeout(function() {
									$(that).typeahead("val", "");
									$(that).typeahead("val", suggestionCopy.textValue);
									$(that).typeahead('open');
								}, 100);
							}
						});
						instance.object.bind('typeahead:asyncrequest', function() {
							 $(instance.selector + '-spinner').show();
						 });
						 instance.object.bind('typeahead:autocompleted :asynccancel typeahead:asyncreceive', function() {
							 $(instance.selector + '-spinner').hide();
						 });
						instance.object.bind("typeahead:active", function() {
							var defaultSearchWidth = 400,
								width = parseInt($(window).width(), 10),
								searchWidth = defaultSearchWidth;
							if (width < searchWidth) {
								searchWidth = width * 9.95 / 10;
								$(instance.selector + '-spinner').css({
									"right": "-110px",
								});
							}
							$(".tt-dataset").css({
								"width": searchWidth - 30,
							});
							setTimeout(function() {
								$(instance.selector + '-spinner').hide();
							}, 100);
						});
						instance.object.bind("typeahead:idle", function() {
							var smallSearch = $("#search-small-input"),
								button = $("#show-small-search-button");
							$("#search-small").hide();
							$("#full-logo").show();
							button.html("<span class='glyphicon glyphicon-search'></span>");
							wisply.isSmallSearchDisplayed = !wisply.isSmallSearchDisplayed;
							$(instance.selector + '-spinner').hide();
						});
					}
					/**
					 * It return the engine which provides the suggestions from the server
					 * @return {object} The engine which provides the results
					 */
					function getSuggestionsEngine() {
						var bloodhound, engine, emptyTemplate, footer, suggestion;
						/**
						 * It returns the bloodhound for the engine
						 * @return {object} The bloodhound
						 * @see https://github.com/twitter/typeahead.js/blob/master/doc/bloodhound.md
						 */
						function getBloodhound() {
							var object = new Bloodhound({
								datumTokenizer: function(d) {
									return Bloodhound.tokenizers.whitespace(d.Title);
								},
								queryTokenizer: Bloodhound.tokenizers.whitespace,
								identify: function(o) {
									return o.ID;
								},
								remote: {
									cacheKey: 'ID',
									url: instance.URL + '%QUERY',
									wildcard: '%QUERY',
								}
							});
							return object;
						}
						suggestion = [
							"<div title='{{ Title }}' style='width:100%'><div class='row'>",
							"<div class='col-lg-1 col-md-1 col-sm-1 col-xs-1 text-center' style='min-width:40px;'>",
							"<div style='width:40px' class='text-center'><img class='search-logo thumbnail' style='margin-bottom:0px' src='{{Icon}}' /></div>",
							"</div>",
							"<div class='col-lg-10 col-md-10 col-sm-10 col-xs-10 search-result'>",
							"<span class='search-title'>{{ cutString Title 40 }}</span><br />",
							"<span class='text-muted bold search-category'><small>{{ Category }}</small></span><br /><span class='text-muted search-description'>{{ cutString Description 90 }}</span></div>",
						"</div></div>",
					].join("\n");
						footer = [
							"<div class='row'><div class='col-lg-12 col-md-12 col-sm-12 col-xs-12'>",
							'<div class="search-footer">',
							"<div><span class='text-primary'><span class='glyphicon glyphicon-search'></span></span> &nbsp;&nbsp;Find more about<strong> {{cutString query 30 }}",
							"</strong></div></div>",
							"</div></div>"
						].join("\n");
						emptyTemplate = [
							'<div class="search-no-results">',
								'<span class="glyphicon glyphicon-inbox"></span> No results available.',
							'</div>'
						].join('\n');
						bloodhound = getBloodhound();
						engine = {
							name: 'Title',
							display: "Title",
							valueKey: "Title",
							source: function(q, sync, async) {
								if (q !== "") {
									bloodhound.search(q, sync, async);
								}
							},
							limit: 7,
							templates: {
								// footer: Handlebars.compile(footer),
								empty: emptyTemplate,
								suggestion: Handlebars.compile(suggestion),
							}
						};
						return engine;
					}
					/**
					 * It is the engine which provides the list for the last searches
					 * It is activated when the input is empty
					 * @return {object} The engine for previous searches
					 */
					function getSearchesEngine() {
						var engine, source;
						source = function() {
							var instanceCopy = instance;
							// keep the code cosistent even if we do not use the event
							/* jshint unused:false */
							return (function cookieEngine(q, sync) {
								if (q === '') {
									var searches = instanceCopy.cookies.getLastSearchQueries();
									sync(searches);
									$(instance.selector + '-spinner').hide();
								}
							});
						}();
						engine = {
							name: 'textValue',
							display: "textValue",
							valueKey: "textValue",
							templates: {
								header: "<div class='search-header'>Previous queries</div>",
								suggestion: Handlebars.compile("<div>{{cutString textValue 50}}</div>"),
							},
							source: source,
						};
						return engine;
					}
					/**
					 * Returns the settings for typeahead object
					 * @return {object} The settings
					 */
					settings = {
						hint: true,
						highlight: false,
						minLength: 0,
					};
					suggestionsEngine = getSuggestionsEngine();
					searchesEngine = getSearchesEngine();
					instance.object = $(instance.selector).typeahead(settings, suggestionsEngine, searchesEngine);
					activateListeners();
				}
				createObject();
				$(".tt-hint").attr('aria-label', 'Search Hint');
			}
		};
	/**
	 * Does nothing
	 * @memberof SearchModule
	 * @class Cookies
	 * @param [SearchModule.Field] field The search field object
	 * @classdesc It manages the operations with cookies
	 */
	var Cookies = function Cookies(field) {
		this.field = field;
		this.cookieName = 'last-searches';
		this.maxNumberOfQueries = 5;
	};
	Cookies.prototype =
		/** @lends SearchModule.Cookies */
		{
			/**
			 * It saves the query in the cookies. The query is saved if it is not the
			 * same as the last one one. If the number of queries exeeds the
			 * allowed one, the last one is removed
			 * @param  {string} newValue The value of the query
			 */
			saveSearchQuery: function(newValue) {
				var list,
					instance = this;
				/**
				 * It changes the value of the cookie
				 * @param  {array} list The array with queries
				 */
				function saveCookie(list) {
					$.cookie(instance.cookieName, JSON.stringify(list), {
						path: '/'
					});
				}
				/**
				 * It checks if the value is not empty and if the last value on the
				 * list is not the same as this
				 * @param  {string}  value The value to be added
				 * @param  {array}  list The current array
				 * @return {Boolean} True if it is valid, false otherwise
				 */
				function isGoodValue(value, list) {
					return (value && value !== "") && ((!list[0]) || (list[0] && value !== list[0].textValue));
				}
				list = this.getLastSearchQueries();
				if (isGoodValue(newValue, list)) {
					list.unshift({
						textValue: newValue,
					});
					if (list.length > this.maxNumberOfQueries) {
						list.pop();
					}
					saveCookie(list);
				}
			},
			/**
			 * It returns the list of saved search queries from cookies
			 * @return {array} The list of saved search queries
			 */
			getLastSearchQueries: function() {
				var string = $.cookie(this.cookieName),
					list = [];
				if (string && string !== "") {
					list = JSON.parse(string);
				}
				return list;
			}
		};
	return {
		Field: Field,
	};
};
$(document).ready(function() {
	"use strict";
	var module = new SearchModule();
	wisply.loadModule("search", module);
});
