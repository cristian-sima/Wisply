<div id="repository-top"></div>
<table class="table">
  <tbody>
    {{range $index, $record := .records}}
    <tr>
      <td>
        <h4>
          {{range $index, $title := $record.Keys.Get "title" }}
          {{ $title }}
          {{ end }}
        </h4>
        {{range $index, $description := $record.Keys.Get "description" }}
        {{ $description }}
        {{ end }}
        <div class="formats">
          {{range $index, $format := $record.Keys.Nice "format" }}
          <span class="label label-default">{{$format}}</span>&nbsp;
          {{ end }}
        </div>
        <div class="creators">
          <span class="text-muted">
            {{range $index, $creator := $record.Keys.Get "creator" }}
            {{ $creator }}
            {{ end }}
          </span>
        </div>
      </td>
    </tr>
    {{ end }}
  </tbody>
</table>

<ul class="pager">
  <li class="previous" ><a href="#">← Older</a></li>
  <li class="next"><a href="#">Newer →</a></li>
</ul>

<div id="repository-bottom"></div>
