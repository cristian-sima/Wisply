<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Account</a></li>
      <li class="active">Searchees</li>
    </ul>
  </div>
  <div class="panel-body">
    <div class="table-responsive">
      <table id="list-accounts" class="table table-striped table-hover ">
        <thead>
          <tr>
            <th class="hidden-xs">#</th>
            <th>Date</th>
            <th>Text</th>
            <th>Accessed the result</th>
          </tr>
        </thead>
        <tbody>
          {{range $index, $search := .searches}}
          {{$query := $search.Query|html}}
          <tr>
            <td class="hidden-xs">{{ $search.ID }}</td>
            <td>{{ $search.GetDate }}</td>
            <td>{{ $query }}</td>
            <td>
              {{ if $search.Accessed }}
              Yes
              {{ else }}
              No
              {{ end }}
            </td>
          </tr>
          {{end }}
        </tbody>
      </table>
    </div>
  </div>
</div>
