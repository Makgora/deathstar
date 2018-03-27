<!DOCTYPE html>
<html>

{{template "base/head.tpl"}}

<body>

{{template "base/navbar.tpl"}}

<div class="container-fluid">
    <div class="row">
        <div id="parkings" class="col-12 col-lg-6">
            {{range .parkings}}
            <div class="parking container">
                <div class="row">
                    <div class="col-2" style="background-color: #dddddd">
                        <img width="100%" class="center" src="static/img/parked/logo_parked.png" alt="Parking photo">
                    </div>
                    <div class="col-10">
                        <div class="row">
                            <div class="col-9">
                                <label>{{.Name}}</label>
                            </div>
                            <div class="col-3">
                                <label>{{.Statut}}</label>
                            </div>
                        </div>
                        <div class="row">
                            <div class="col-9">
                                <label class="text-muted">{{.Address}}</label>
                            </div>
                            <div class="col-3">
                                <label>{{.FreeSpaces}} places</label>
                            </div>
                        </div>
                        <div class="row">
                            <div class="col-9">
                                <button type="button" class="btn btn-outline-secondary"disabled>Réserver</button>
                            </div>
                            <div class="col-3">
                                {{if not .Address}}
                                    <button type="button" class="btn btn-outline-secondary"disabled>Go</button>
                                {{else}}
                                    <a href="https://www.google.fr/maps/dir//{{.Address}}" target="_blank" type="button" class="btn btn-outline-success btn-success">Go</a>
                                {{end}}
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            {{end}}
        </div>
        <div class="d-none d-lg-block col-lg-6">
            <div id="map">
            <script>
                //TODO V
                var cityLocation = {lat: {{.cityLocation.Lat}}, lng: {{.cityLocation.Lng}}};
                var beaches = [];
                {{range .parkings}}
                        beaches.push(["{{.Name}}", {{.Location.Lat}}, {{.Location.Lng}}]);
                {{end}}
                console.log(beaches);
                initMap()
            </script>
            </div>
        </div>
    </div>
    <div class="row">
        <div class="col-12 col-lg-6">
            <label id="update" class="text-muted font-italic">
                Updated {{.lastUpdate}} sec ago
            </label>
        </div>
    </div>
</div>

{{template "base/footer.tpl"}}

{{template "base/scripts.tpl"}}

<script async defer src="https://maps.googleapis.com/maps/api/js?key=AIzaSyBVFpJA9GjD8dqEu0KHVwNjQTZ6XwrnCFQ&callback=initMap"></script>
<script src="static/js/map.js"></script>

</body>

</html>