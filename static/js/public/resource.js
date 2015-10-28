/* global $, wisply */
var login;

/**
* @file Encapsulates the functionality for displaying a public resource
* @author Cristian Sima
*/

/**
 * Requires nothing
* @namespace PublicResource
*/
var PublicResource = function () {
  'use strict';
  var page;
  /**
  * Encapsulates the functionality for public resources
  * @class Page
  * @memberof LoginModule
  * @classdesc It represents the page
  * @param {object} o The object with the repository and resource objects
  */
  var Page = function Page(o) {
    this.iframe = new IFrame(this);
    this.repository = o.repository;
    this.resource = o.resource;
  };
  Page.prototype =
  /** @lends PublicResource.Page */
  {
    /**
    * It loads the listeners
    */
    init : function () {
      this.iframe.init();
    },
  };

  /**
  * @class IFrame
  * @memberof LoginModule
  * @classdesc It represents the IFrame
  * @param {PublicResource.Page} page The reference to the page object
  */
  var IFrame = function IFrame(paga) {
    this.element = $("#the-iframe");
    this.page = page;
  };
  IFrame.prototype =
  /** @lends PublicResource.IFrame */
  {
    /**
    * It shows the loading
    */
    init : function () {
      this._showLoading();
    },
    /**
     * It shows the loading image
     */
    _showLoading: function () {
      //src="/repository/{{ .repository.ID }}/resource/{{ .resource.ID }}/content"
      var html = "<div style='text-align:center;margin-top:50px; margin-bottom:50px;' >" + wisply.getLoadingImage("medium") + "<br />Loading content...</div>";
      this.element.contents().find("body").html(html);
    }
  };

  /**
   * Creates the page object and stores it in the page variable
   * @param  {object} object The object which contains the repository and the
   * resource
   * @return {PublicResource.Page} The page
   */
  function init(object) {
    page = new Page(object);
    page.init();
    return page;
  }
  return {
      init: init,
      Page: Page
  };
};
$(document).ready(function() {
  "use strict";
  wisply.loadModule("public-resource", new PublicResource());
});
