/* global $, wisply */

/**
* @file Encapsulates the functionality for only the administrators on all the pages
* @author Cristian Sima
*/

/**
* @namespace Administration
*/
var Administration = function () {
  'use strict';

  /**
  * The constructor calls the init method
  * @class Admin
  * @memberof Administration
  * @classdesc It manages the operations of administrators on all the pages
  */
  var Admin = function Admin() {
  };

  Admin.prototype =
  /** @lends Administration.Admin */
  {
    /**
    * It activates the button for displaying the sidebar
    */
    init: function () {
      $('[data-toggle=offcanvas], #close-sidebar-admin').click(function() {
        $('.row-offcanvas').toggleClass('active');
        $('html, body').animate({ scrollTop: 0 }, 'fast');
      });
    }
  };
  return {
    Admin: Admin
  };
};
$(document).ready(function() {
  "use strict";
  var module = new Administration();
  wisply.admin = new module.Admin();
  wisply.admin.init();
});
