<html>
<!DOCTYPE html>
{{template "base/head.tpl"}}
<body>
{{template "base/navbar.tpl"}}

<div class="container mgt">
    <h1 class="display-1 text-center">Nous contacter ?</h1>
    <p class="lead text-center" style="margin-top: 50px">
        N'hésitez pas à nous donner un feedback sur l'application, nous acceptons tout encouragement ou subvention gracieuse ! =P
    </p>
</div>

<div class="container mgt">
    <div class="form-group row">
        <div class="col-md-2 offset-md-3 col-lg-2 offset-lg-3">
            <label for="name-input" class="col-form-label">Name</label>
        </div>
        <div class="col-md-5 col-lg-4">
            <input class="form-control" type="text" placeholder="Your name" id="name-input">
        </div>
    </div>
    <div class="form-group row">
        <div class="col-md-2 offset-md-3 col-lg-2 offset-lg-3">
            <label for="email-input" class="col-form-label">Email</label>
        </div>
        <div class="col-md-5 col-lg-4">
            <input class="form-control" type="email" placeholder="Your email" id="email-input">
        </div>
    </div>
    <div class="form-group row">
        <div class="col-md-2 offset-md-3 col-lg-2 offset-lg-3">
            <label for="msg-input" class="col-form-label">Your message</label>
        </div>
        <div class="col-md-5 col-lg-4">
            <textarea class="form-control" type="text" placeholder="Please enter your message here..." id="msg-input"></textarea>
        </div>
    </div>
    <div class="text-center">
        <button type="button" class="btn btn-sm" style="margin-top: 20px">Submit</button>
    </div>
</div>

{{template "base/footer.tpl"}}
{{template "base/scripts.tpl"}}
</body>
</html>