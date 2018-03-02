{{define "footer"}}
<footer>
    <div class="author">
        Official website:
        <a href="http://{{.Website}}">{{.Website}}</a>
        <br>
        Contact me:
        <a class="email" href="mailto:{{.Email}}">{{.Email}}</a>
    </div>
</footer>
<script src="/static/js/reload.min.js"></script>
</body>
</html>
{{end}}