/* global $, wisply, Handlebars */

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
      minLength: 2,
    },
    {
      name: 'states',
      source: searchAnything,
  		display: 'Title',
      templates: {
      pending :  [
          "<div class='text-center empty-message' >",
          wisply.getLoadingImage("small"),
          "<br /></div>"
        ].join('\n'),
      empty: [
        '<div class="empty-message">',
          '<span class="glyphicon glyphicon-inbox"></span> It seems there is nothing like that',
        '</div>'
      ].join('\n'),
      header: [
				"<h4 class='search-header league-name'>",
				"Institutions",
				"</h4>",
				"<hr />",
			].join("\n"),
    	suggestion: Handlebars.compile([
				"<div>",
				"<strong>{{Title}}</strong>",
					"<div class='row'>",
						"<div class='col-lg-2 col-md-2 col-sm-2'>",
						"<img class='search-logo' src='{{Data.LogoURL}}' />",
						"</div>",
						"<div class='col-lg-10 col-md-10 col-sm-10 text-muted smaller'>{{ trimS Data.Description 0 150 }}</div>",
					"</div>",
				"</div>",
			].join("\n")),
      }
    });
		this.object.bind('typeahead:select', function(ev, suggestion) {
		  window.location = suggestion.URL;
		});
	};
	Handlebars.registerHelper('trimS', function(passedString, start, length ){
		var mlength = length,preS='',tailS='';
		if(start>0 && passedString.length>3){
		    preS= '...';
		    mlength = length -3;
		}
		if(passedString.length>(start + length )){
		    tailS = '...';
		    mlength = mlength -3;
		}
		var theString = preS + passedString.substr(start, mlength) + tailS;
		return new Handlebars.SafeString(theString);
	});
	Field.prototype =
		/** @lends SearchModule.Field */
		{

		};
	return {
		Field: Field,
	};
};
$(document).ready(function() {
	"use strict";
	searchModule = new SearchModule();
});
