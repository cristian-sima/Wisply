<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li class="active">Accounts</li>
    </ul>
  </div>
  <div class="panel-body">
    <div class="table-responsive">
      <table class="table table-striped table-hover ">
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
          {{range $index, $element := .accounts}}
          {{$safe := $element.Email|html}}
          <tr>
            <td class="hidden-xs">{{ $element.ID |html }}</td>
            <td>{{ $element.Name |html }}</td>
            <td><a href="mailto:{{ $safe }}">{{ $element.Email |html }}</a></td>
            <td>
              {{ if $element.IsAdministrator }}
              <span class="label label-info">Administrator</span>
              {{ else }}
              <span class="label label-default">User</span>
              {{ end }}
            </td>
            <td>
              <a href="/admin/accounts/modify/{{$element.ID}}">Modify</a>
            </td>
            <td>
              <a class="deleteAccountButton" data-id="{{$element.ID}}" data-name="{{$safe}}" href="/">Delete</a>
            </td>
          </tr>
          {{end }}
        </tbody>
      </table>
    </div>
  </div>
</div>
<script src="/static/js/static/account/list.js"></script>
