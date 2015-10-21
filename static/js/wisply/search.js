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
	// Maximum allowed characters for description
	var maxAllowedCharForDesc = 90;
	/**
	 * Does nothing
	 * @memberof Field
	 * @class Wikier
	 * @param [string] select The selector for the search
	 * @classdesc It gets the elements
	 */
	var Field = function Field(selector) {
		var searchAnything = new Bloodhound({
			datumTokenizer: Bloodhound.tokenizers.obj.whitespace('Title'),
			queryTokenizer: Bloodhound.tokenizers.whitespace,
			remote: {
				url: '/api/search/anything/%QUERY',
				wildcard: '%QUERY'
			}
		});
		this.object = $(selector).typeahead({
			hint: true,
			highlight: false,
			minLength: 1,
		}, {
			name: 'states',
			source: searchAnything,
			display: 'Title',
			templates: {
				pending: [
          "<div class='text-center empty-message' >",
          wisply.getLoadingImage("small"),
          "<br /></div>"
        ].join('\n'),
				empty: [
        '<div class="empty-message">',
          '<span class="glyphicon glyphicon-inbox"></span> No results',
        '</div>'
      ].join('\n'),
				header: [
				'<div class="search-header">',
  			"<span class='h6' class='league-name'>",
				"Institutions",
				"</span></div>",
			].join("\n"),
				suggestion: Handlebars.compile([
					"<div style='width:100%'><div class='row'>",
					"<div class='col-lg-2 col-md-2 col-sm-2 col-xs-2'>",
					"<img class='search-logo' src='{{Icon}}' />",
					"</div>",
					"<div class='col-lg-10 col-md-10 col-sm-10 col-xs-10 search-result'>",
					"<span class='search-title'>{{Title}}</span><br />",
					"<span class='text-muted search-description'>{{ cutString Description " + maxAllowedCharForDesc + " }}</span></div>",
				"</div></div>",
			].join("\n")),
			}
		}, {
			name: 'states',
			source: searchAnything,
			display: 'Title',
			templates: {
				pending: [
          "<div class='text-center empty-message' >",
          wisply.getLoadingImage("small"),
          "<br /></div>"
        ].join('\n'),
				empty: [
        '<div class="empty-message">',
          '<span class="glyphicon glyphicon-inbox"></span> No results',
        '</div>'
      ].join('\n'),
				header: [
				'<div class="search-header">',
  			"<span class='h6' class='league-name'>",
				"Institutions",
				"</span></div>",
			].join("\n"),
				suggestion: Handlebars.compile([
					"<div style='width:100%'><div class='row'>",
					"<div class='col-lg-2 col-md-2 col-sm-2 col-xs-2'>",
					"<img class='search-logo' src='{{Icon}}' />",
					"</div>",
					"<div class='col-lg-10 col-md-10 col-sm-10 col-xs-10 search-result'>",
					"<span class='search-title'>{{Title}}</span><br />",
					"<span class='text-muted search-description'>{{ cutString Description " + maxAllowedCharForDesc + " }}</span></div>",
				"</div></div>",
			].join("\n")),
			}
		});
    // keep the code cosistent even if we do not use the event
    /* jshint unused:false */
		this.object.bind('typeahead:select', function(event, suggestion) {
			window.location = suggestion.URL;
		});
		this.object.bind("typeahead:active", function(){
				var defaultSearchWidth = 400,
					width = parseInt($(window).width(), 10),
					searchWidth = defaultSearchWidth;
				if(width < searchWidth) {
					searchWidth = width *9.95/10;
				}
				$(".tt-menu").css({"width": searchWidth - 30,});
		});
		this.object.bind("typeahead:idle", function() {
				var smallSearch = $("#search-small-input"),
				button = $("#show-small-search-button");
				$("#search-small").hide();
				$("#full-logo").show();
				button.html("<span class='glyphicon glyphicon-search'></span>");
				wisply.isSmallSearchDisplayed = !wisply.isSmallSearchDisplayed;
		});
	};
	Field.prototype =
		/** @lends SearchModule.Field */
		{};
	return {
		Field: Field,
	};
};
$(document).ready(function() {
	"use strict";
	searchModule = new SearchModule();
});
