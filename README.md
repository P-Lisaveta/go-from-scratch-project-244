# Вычислитель отличий на Go

[![CI](https://github.com/P-Lisaveta/go-from-scratch-project-244/actions/workflows/ci.yml/badge.svg)](https://github.com/P-Lisaveta/go-from-scratch-project-244/actions/workflows/ci.yml)
[![hexlet-check](https://github.com/P-Lisaveta/go-from-scratch-project-244/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/P-Lisaveta/go-from-scratch-project-244/actions/workflows/hexlet-check.yml)

Консольная утилита для сравнения вложенных структур (JSON, YAML)

Учебный проект Хекслета: https://ru.hexlet.io/programs/go-from-scratch

## Стек

- Go
- [urfave/cli](https://github.com/urfave/cli)
- [gopkg.in/yaml.v3](https://github.com/go-yaml/yaml)

## Установка

```bash
git clone https://github.com/P-Lisaveta/go-from-scratch-project-244.git
cd go-from-scratch-project-244
make build
```

## Использование

```bash
./bin/gendiff testdata/fixture/file1.json testdata/fixture/file2.json
./bin/gendiff testdata/fixture/file1.yml testdata/fixture/file2.yml
./bin/gendiff testdata/fixture/nested1.json testdata/fixture/nested2.json
./bin/gendiff testdata/fixture/nested1.yml testdata/fixture/nested2.yml
./bin/gendiff --format stylish testdata/fixture/file1.json testdata/fixture/file2.json
./bin/gendiff -f stylish testdata/fixture/file1.yml testdata/fixture/file2.yml
./bin/gendiff --format plain testdata/fixture/nested1.json testdata/fixture/nested2.json
./bin/gendiff -f plain testdata/fixture/nested1.yml testdata/fixture/nested2.yml
./bin/gendiff --format json testdata/fixture/nested1.json testdata/fixture/nested2.json
./bin/gendiff -f json testdata/fixture/nested1.yml testdata/fixture/nested2.yml
```

## Разработка

```bash
make test
make lint
make test-coverage
```

Минимальный порог покрытия задается переменной `COVERAGE_MIN` и по умолчанию равен `80`:

```bash
COVERAGE_MIN=95 make test-coverage
```

## Демонстрация

Аскинемы с примерами сравнения конфигураций:

- [Плоский JSON](docs/gendiff-flat-json.cast)
- [Плоский YAML](docs/gendiff-flat-yaml.cast)
- [Вложенный JSON, формат stylish](docs/gendiff-nested-json.cast)
- [Вложенный JSON, формат plain](docs/gendiff-plain-json.cast)
- [Вложенный JSON, формат json](docs/gendiff-json.cast)

Просмотреть локальную запись можно командой:

```bash
asciinema play docs/gendiff-flat-json.cast
asciinema play docs/gendiff-flat-yaml.cast
asciinema play docs/gendiff-nested-json.cast
asciinema play docs/gendiff-plain-json.cast
asciinema play docs/gendiff-json.cast
```

---

<details>
<summary>Автоматические тесты Хекслета</summary>

Тесты запускаются на каждый коммит. За запуск отвечает файл `.github/workflows/hexlet-check.yml` — не удаляйте и не переименовывайте ни его, ни репозиторий.

</details>

## О Хекслете

[Хекслет](https://ru.hexlet.io/) — школа программирования: авторские программы обучения с практикой, поддержкой наставников и реальными проектами, которые остаются в резюме. Этот репозиторий — один из таких проектов.
