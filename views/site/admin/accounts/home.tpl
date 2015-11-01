<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li class="active">Accounts</li>
    </ul>
  </div>
  <div class="panel-body">
    <div class="table-responsive">
      <table id="list-accounts" class="table table-striped table-hover ">
        <thead>
          <tr>
            <th class="hidden-xs">Id</th>
            <th>Name</th>
            <th>E-mail</th>
            <th>Type</th>
            <th>Modify</th>
            <th>Delete</th>
          </tr>
        </thead>
        <tbody>
          {{range $index, $account := .accounts}}
          <tr>
            <td class="hidden-xs">{{ $account.ID |html }}</td>
            <td>{{ $account.Name |html }}</td>
            <td><a href="mailto:{{ $account.Email }}">{{ $account.Email }}</a></td>
            <td>
              {{ if $account.IsAdministrator }}
              <span class="label label-info">Administrator</span>
              {{ else }}
              <span class="label label-default">User</span>
              {{ end }}
            </td>
            <td>
              <a href="/admin/accounts/{{$account.ID}}/modify">Modify</a>
            </td>
            <td>
              <a class="deleteAccountButton" data-id="{{$account.ID}}" data-name="{{$account.Name}}" href="#">Delete</a>
            </td>
          </tr>
          {{end }}
        </tbody>
      </table>
    </div>
  </div>
</div>
<script src="/static/js/admin/account/list.js"></script>
<script>
$(document).ready(function(){
    var module = wisply.getModule("admin-accounts"),
      list = new module.List();
      list.init();
});
</script>
