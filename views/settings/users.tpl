{{template "base/base.tpl" .}}
{{define "head"}}
    <title>{{.title}}</title>
{{end}}
{{define "body"}}
    <div class="page-header d-print-none m-2">
        <div class="container-xl">
            <div class="row g-2 align-items-center">
                <div class="col-auto ms-auto d-print-none">
                    <div class="d-flex">
                        <a href="#" class="btn btn-primary">
                            <svg xmlns="http://www.w3.org/2000/svg" class="icon" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"></path><path d="M12 5l0 14"></path><path d="M5 12l14 0"></path></svg>
                            Создать пользователя
                        </a>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <div class="container-xl">
        <div class="row row-cards">
            <div class="col-lg-12">
                <div class="card">
                    <div class="table-responsive">
                        <table class="table table-vcenter card-table">
                            <thead>
                            <tr>
                                <th>Логин</th>
                                <th>Последний вход</th>
                                <th>Роль</th>
                                <th>Статус</th>
                                <th class="w-1"></th>
                            </tr>
                            </thead>
                            <tbody>
                            <tr>
                                <td>admin</td>
                                <td class="text-secondary">2024-03-27 13:05:22</td>
                                <td class="text-secondary"><a href="#" class="text-reset">Админ</a></td>
                                <td class="text-secondary">Включен</td>
                                <td><a href="#">Редактировать</a></td>
                            </tr>
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>

        </div>
    </div>
{{end}}