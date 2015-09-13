{{ define "javascript" }}

<script src="/static/3rd_party/product/paper/jquery-1.10.2.min.js"></script>
<script src="/static/3rd_party/product/paper/bootstrap.min.js"></script>
<script src="/static/3rd_party/product/paper/bootswatch.js"></script>
<script src="/static/3rd_party/product/paper/paper.js"></script>

<script src="/static/3rd_party/others/js/bootbox.min.js"></script>
<script src="/static/3rd_party/others/js/base64_decode.js"></script>
<script src="/static/3rd_party/others/js/jquery.cookie.js"></script>
<script src="/static/3rd_party/others/js/jquery.hotkeys.js"></script>

<script src="/static/js/wisply/wisply.js"></script>

{{ if .accountConnected }}
<script src="/static/js/wisply/Connection.js"></script>
{{ end }}

{{ end }}
