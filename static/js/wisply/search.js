/* global $, wisply */

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
      datumTokenizer: Bloodhound.tokenizers.obj.whitespace('value'),
      queryTokenizer: Bloodhound.tokenizers.whitespace,
      remote: {
        url: '/api/search/anything/%QUERY',
        wildcard: '%QUERY'
      }
    });

    $(selector).typeahead({
      hint: true,
      highlight: true,
      minLength: 1,
    },
    {
      name: 'states',
      source: searchAnything,
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
      header: '<h3 class="league-name">Something</h3>'
      }
    });

	};
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
