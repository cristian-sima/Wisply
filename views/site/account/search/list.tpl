<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/account">Account</a></li>
      <li class="active">Search History</li>
    </ul>
  </div>
  <div class="panel-body">
  {{ if eq (.searches | len ) 0 }}
  There is no search available
  {{ else }}
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
              <span class="label label-success">Yes</span>
              {{ else }}
              <span class="label label-default">No</span>
              {{ end }}
            </td>
          </tr>
          {{end }}
        </tbody>
      </table>
    </div>
    <a class="btn btn-danger" id="clearSearchHistory">Clear my search history</a>
    {{ end }}
  </div>
</div>
<script src="/static/js/account/search/list.js"></script>
