/* global jQuery, wisply */

(function ($) {
  'use strict';

  /**
  * It manages the general functionality for administrators
  */
  function Admin() {
    this.init();
  }

  Admin.prototype = {

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

  /**
  * It is called when the page has been loaded. It creates the Admin object
  */
  function initPage() {
    wisply.admin = new Admin();
  }
  $(document).ready(initPage);
}(jQuery, wisply));
