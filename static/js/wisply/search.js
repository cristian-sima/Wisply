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
	// register a new function for the JavaScript template library
	Handlebars.registerHelper('cutString', function(theString, numberOfAllowedCharacters) {
		if (theString.length > numberOfAllowedCharacters) {
			var shortString = theString.substr(0, numberOfAllowedCharacters);
			theString = shortString + "...";
		}
		return new Handlebars.SafeString(theString);
	});
	/**
	 * Does nothing
	 * @memberof Field
	 * @class Wikier
	 * @param [string] select The selector for the search
	 * @classdesc It gets the elements
	 */
	var Field = function Field(o) {

		this.selector = o.selector;
		this.URL = o.URL;
		this.saveSearches = o.saveSearches;

		var instance = this,
			engine;

		engine = new Bloodhound({
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
					rateLimitWait: 1000,
				}
			});

		var suggestion = [
					"<div style='width:100%'><div class='row'>",
					"<div class='col-lg-1 col-md-1 col-sm-1 col-xs-1 text-center' style='min-width:40px;'>",
					"<div style='width:40px' class='text-center'><img class='search-logo thumbnail' style='margin-bottom:0px' src='{{Icon}}' /></div>",
					"</div>",
					"<div class='col-lg-10 col-md-10 col-sm-10 col-xs-10 search-result'>",
					"<span class='search-title'>{{ cutString Title 40 }}</span><br />",
					"<span class='text-muted bold search-category'><small>{{ Category }}</small></span><br /><span class='text-muted search-description'>{{ cutString Description 110 }}</span></div>",
				"</div></div>",
			].join("\n");
		var footer = [
					"<div class='row'><div class='col-lg-12 col-md-12 col-sm-12 col-xs-12'>",
					'<div class="search-footer">',
					"<span><span class='text-primary'><span class='glyphicon glyphicon-search'></span></span> &nbsp;&nbsp;Find more about<strong> {{cutString query 30 }}",
					"</strong></div>",
					"</div></div>"
				].join("\n");
		var emptyTemplate = [
					'<div class="search-no-results">',
						'<span class="glyphicon glyphicon-inbox"></span> No results available.',
					'</div>'
				].join('\n');
		var resultEngine = {
			name: 'Title',
			display: "Title",
			valueKey: "Title",
			source: function(q, sync, async) {
				if (q !== "") {
					engine.search(q, sync, async);
				}
			},
			limit: 7,
			templates: {
				footer: Handlebars.compile(footer),
				empty: emptyTemplate,
				suggestion: Handlebars.compile(suggestion),
			}
		};
		var cookieEngine = {
			templates: {
				header: "<div class='search-header'>Last searches</div>",
			},
			source: function() {
				var instanceCopy = instance;
				return (function cookieEngine(q, sync, async) {
					if (q === '') {
						var searches = instanceCopy.getLastSearchQueries(),
							a = instanceCopy.getLastSearchQueries();
						sync(a);
					}
				});
			}(),
		};
		this.object = $(instance.selector).typeahead({
			hint: true,
			highlight: false,
			minLength: 0,
		}, resultEngine, cookieEngine);
		// keep the code cosistent even if we do not use the event
		/* jshint unused:false */
		this.object.bind('typeahead:select', function(event, suggestion) {
			if (suggestion.Category) {
				if(instance.saveSearches) {
					instance.saveSearchQuery(suggestion.Title);
				}
				if(wisply.connection) {
					var	requestObject = {
							url: "/api/search/save/" + suggestion.Title,
							success: function() {
								window.location = suggestion.URL;
							},
						};
					wisply.executePostAjax(requestObject);
				} else {
						window.location = suggestion.URL;
				}
			} else {
				var t = this,
					suggestionCopy = suggestion;
				$(this).typeahead('close');
				setTimeout(function() {
					$(t).typeahead("val", "");
					$(t).typeahead("val", suggestionCopy);
					$(t).typeahead('open');
				}, 100);
			}
		});
		this.object.bind("typeahead:active", function() {
			var defaultSearchWidth = 400,
				width = parseInt($(window).width(), 10),
				searchWidth = defaultSearchWidth;
			if (width < searchWidth) {
				searchWidth = width * 9.95 / 10;
			}
			$(".tt-menu").css({
				"width": searchWidth - 30,
			});
		});
		this.object.bind("typeahead:idle", function() {
			var smallSearch = $("#search-small-input"),
				button = $("#show-small-search-button");
			$("#search-small").hide();
			$("#full-logo").show();
			button.html("<span class='glyphicon glyphicon-search'></span>");
			wisply.isSmallSearchDisplayed = !wisply.isSmallSearchDisplayed;
		});
		this.cookieName = 'last-searches';
	};
	Field.prototype =
		/** @lends SearchModule.Field */
		{
			saveSearchQuery: function(newValue) {
				var oldList,
					allowedSearches = 5,
					newList;
				oldList = this.getLastSearchQueries();
				newList = oldList;
				if (!oldList) {
					newList = [];
				}
				if (newValue && newValue !== "") {
					// if the same as last one
					if (!oldList || (oldList && newValue !== oldList[0])) {
						newList.unshift(newValue);
						if (newList.length > allowedSearches) {
							newList.pop();
						}
						$.cookie(this.cookieName, JSON.stringify(newList), {
							path: '/'
						});
					}
				}
			},
			getLastSearchValue: function() {
				var listOfSearches = this.getLastSearchQueries(),
					last;
				if (!listOfSearches) {
					return "";
				}
				last = listOfSearches[0];
				if (!last) {
					return "";
				}
				return last;
			},
			getLastSearchQueries: function() {
				var string =  $.cookie(this.cookieName);
				if(string === "" || !string) {
					return [];
				}
				return JSON.parse(string);
			}
		};
	return {
		Field: Field,
	};
};
$(document).ready(function() {
	"use strict";
	searchModule = new SearchModule();
});
