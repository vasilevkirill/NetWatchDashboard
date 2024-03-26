<!doctype html>
<html lang="ru">
<head>
    <meta charset="utf-8"/>
    <meta name="viewport" content="width=device-width, initial-scale=1, viewport-fit=cover"/>
    <meta http-equiv="X-UA-Compatible" content="ie=edge"/>
    <title>Авторизация</title>
    <link href="/static/tabler.min.css" rel="stylesheet"/>
</head>
<body  class=" d-flex flex-column bg-white">
<div class="row g-0 flex-fill">
    <div class="col-12 col-lg-6 col-xl-4 border-top-wide border-primary d-flex flex-column justify-content-center">
        <div class="container container-tight my-5 px-lg-5">
            <div class="text-center mb-4 ">
                <a href="/" class="navbar-brand navbar-brand-autodark"><img src="/static/66016cf86cb97.png" class="rounded" height="128" alt=""></a>
            </div>
            <h2 class="h3 text-center mb-3">Авторизация</h2>
            {{if .error}}
                <div class="alert alert-danger" role="alert">
                    {{.error}}
                </div>
            {{end}}
            <form action="/login" method="post">
                <div class="mb-3">
                    <label class="form-label">Имя пользователя</label>
                    <input name="username" type="text" class="form-control" placeholder="username" required>
                </div>
                <div class="mb-2">
                    <label class="form-label">
                        Пароль
                    </label>
                    <div class="input-group input-group-flat">
                        <input name="password" type="password" class="form-control"  placeholder="password" required>
                    </div>
                </div>
                <div class="mb-2">
                    <label class="form-check">
                        <input name="saveme" type="checkbox" class="form-check-input"/>
                        <span class="form-check-label">Запоминить меня</span>
                    </label>
                </div>
                <div class="form-footer">
                    <button type="submit" class="btn btn-primary w-100">Вход</button>
                </div>
            </form>
        </div>
    </div>
    <div class="col-12 col-lg-6 col-xl-8 d-none d-lg-block">
        <div class="bg-cover h-100 min-vh-100" style="background-image: url('/static/660164f777061.png')"></div>
    </div>
</div>
<script src="/static/tabler.min.js" defer></script>
</body>
</html>