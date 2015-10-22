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
	// Array Remove - By John Resig (MIT Licensed)
	Array.prototype.remove = function(from, to) {
	  var rest = this.slice((to || from) + 1 || this.length);
	  this.length = from < 0 ? this.length + from : from;
	  return this.push.apply(this, rest);
	};
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
		var engine = new Bloodhound({
			datumTokenizer: Bloodhound.tokenizers.obj.whitespace('ID'),
			queryTokenizer: Bloodhound.tokenizers.whitespace,
			remote: {
				url: '/api/search/anything/%QUERY',
				wildcard: '%QUERY'
			}
		});


		function getTemplate(type) {
			function getSuggestion() {
				return [
			"<div style='width:100%'><div class='row'>",
			"<div class='col-lg-1 col-md-1 col-sm-1 col-xs-1 text-center' style='min-width:40px;'>",
			"<div style='width:40px' class='text-center'><img class='search-logo thumbnail' style='margin-bottom:0px' src='{{Icon}}' /></div>",
			"</div>",
			"<div class='col-lg-10 col-md-10 col-sm-10 col-xs-10 search-result'>",
			"<span class='search-title'>{{ cutString Title " + maxAllowedCharForDesc + " }}</span><br />",
			"<span class='text-muted bold'><small>{{ Category }}</small></span><br /><span class='text-muted search-description'>{{ cutString Description " + maxAllowedCharForDesc + " }}</span></div>",
		"</div></div>",
	].join("\n");
			}
			var searchMoreOption = [
			"<div class='row'><div class='col-lg-12 col-md-12 col-sm-12 col-xs-12'>",
			'<div class="search-footer">',
			"<span class='h6'><span class='text-primary'><span class='glyphicon glyphicon-search'></span></span> &nbsp;&nbsp;Find more about<strong> {{query}}",
			"</strong></div>",
			"</div></div>"
		].join("\n");
			return {
				name: 'Title',
				source: engine,
				display: "Title",
				templates: {
				footer: Handlebars.compile(searchMoreOption),
				empty: [
        '<div class="empty-message">',
          '<span class="glyphicon glyphicon-inbox"></span> No results available.',
        '</div>'
      ].join('\n'),
				suggestion: Handlebars.compile(getSuggestion(type)),
			}
			};
		}
		this.object = $(selector).typeahead({
			hint: true,
			highlight: false,
			minLength: 1,
		}, getTemplate("Institution"));
		// keep the code cosistent even if we do not use the event
		/* jshint unused:false */
		this.object.bind('typeahead:select', function(event, suggestion) {
			window.location = suggestion.URL;
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
