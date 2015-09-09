$(document).ready(function() {
  initPage();
});

function initPage () {
  loadListeners();
}

function loadListeners() {
  $("#menu-logout-button").click(logoutUser);
}

function logoutUser (event) {
  setElementLoading($("#menu-logout-button"), "small");
  event.preventDefault();
  executePostAjax({
    "url" : '/auth/logout',
    dataType: "text",
    'method': "POST",
    "type" : "POST",
    "success": function() {
      showSuccessMessage("You have been disconnected! Refreshing page...");
      reloadPage();
      $("#menu-logout-button").html("");
    },
    "error" : function() {
      showErrorMessage("There was a problem with your request!");
    }
  });

}

function executePostAjax(arguments) {
  if(!arguments.data) {
    arguments.data = {};
  }
  var xsrf,
  xsrflist
  xsrf = $.cookie("_xsrf");
  xsrflist = xsrf.split("|");
  arguments.data._xsrf = base64_decode(xsrflist[0]);
  $.ajax(arguments);
}

function showSuccessMessage (message) {
  showMessage("<div class='text-success'>Success</div>", message);
}
function showErrorMessage (message) {
  showMessage("<div class='bg-warning>Sorry</div>", message);
}
function showMessage(title, content) {
  bootbox.dialog({
    title: title,
    message: content
  });
}
function reloadPage() {
  setTimeout(function() {
    location.reload();
  }, 2000);
}
function setElementLoading(element, size) {
  switch(size) {
    case "small":
    px = 20;
    break;
    case "medium":
    px = 55;
    break;
    case "big":
    px = 110
    break;
  }
  element.html("<img src='/static/img/main/load.gif' style='height: " + px + "px; width: " + px + "px' />");
}
