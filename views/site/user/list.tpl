<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li class="active">Users</li>
    </ul>
  </div>
  <div class="panel-body">
    {{ if .anything }}
    <div class="table-responsive">
      <table class="table table-striped table-hover ">
        <thead>
          <tr>
            <th class="hidden-xs">Id</th>
            <th>Username</th>
            <th>E-mail</th>
            <th>Type</th>
            <th style="text-align:left">Options</th>
          </tr>
        </thead>
        <tbody>
          {{range $index, $element := .users}}
          {{$safe := $element.Email|html}}
          <tr>
            <td class="hidden-xs">{{ $element.Id |html }}</td>
            <td>{{ $element.Username |html }}</td>
            <td><a href="mailto:{{ $safe }}">{{ $element.Email |html }}</a></td>
            <td>
              {{ if $element.Administrator }}
              <span class="label label-info">Administrator</span>
              {{ else }}
              <span class="label label-default">User</span>
              {{ end }}
            </td>
            <td>
              <div class="btn-group"  >
                <a href="/"  class="btn btn-link dropdown-toggle btn-sm" data-toggle="dropdown"><span class="caret"></span></a>
                <ul class="dropdown-menu" >
                  <li><a href="/admin/users/modify/{{$element.Id}}">Modify</a></li>
                  <li><a class="deleteUserButton" data-id="{{$element.Id}}" data-name="{{$safe}}" href="/">Delete</a></li>
                </ul>
              </div>
            </td>
          </tr>
          {{end }}
        </tbody>
      </table>
    </div>
    {{ else }}
    There are no users ... :(
    {{ end }}
  </div>
</div>
<script>
/* global bootbox */
$(document).ready(function () {
  $(".deleteUserButton").click(confirmDelete);
});

function confirmDelete(e) {
  e.preventDefault();

  var instance = $(this),
  id = instance.data("id"),
  name = instance.data("name");
  showConfirmationDialog(id, name);
}

function showConfirmationDialog(id, name) {
  bootbox.dialog({
    title: "Confirmation",
    message: "The user <b>" + name + "</b> will be permanently removed. Are you sure?",
    onEscape: true,
    buttons: {
      cancel: {
        label: "Cancel",
        className: "btn-default",
        callback: function() {
          this.modal('hide');
        }
      },
      main: {
        label: "Delete",
        className: "btn-primary",
        callback: function() {
          deleteSource(id);
        }
      }
    }
  });
}
function deleteSource(id) {
  executePostAjax({
    "url" : '/admin/users/delete/' + id,
    dataType: "text",
    'method': "POST",
    "type" : "POST",
    "success": function() {
      showSuccessMessage("The user has been removed! Refreshing page...");
      reloadPage();
    },
    "error" : function() {
      showErrorMessage("There was a problem with your request!");
    }
  })
}
</script>
