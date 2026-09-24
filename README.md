# Вычислитель отличий на Go

[![hexlet-check](https://github.com/P-Lisaveta/go-from-scratch-project-244/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/P-Lisaveta/go-from-scratch-project-244/actions/workflows/hexlet-check.yml)

Консольная утилита для сравнения вложенных структур (JSON, YAML)

Учебный проект Хекслета: https://ru.hexlet.io/programs/go-from-scratch

## Стек

- Go
- [urfave/cli](https://github.com/urfave/cli)

## Установка

```bash
git clone https://github.com/P-Lisaveta/go-from-scratch-project-244.git
cd go-from-scratch-project-244
make build
```

## Использование

```bash
./bin/gendiff tests/fixtures/file1.json tests/fixtures/file2.json
./bin/gendiff --format stylish tests/fixtures/file1.json tests/fixtures/file2.json
./bin/gendiff -f stylish tests/fixtures/file1.json tests/fixtures/file2.json
```

## Демонстрация

Аскинема с примером сравнения плоских JSON-файлов:

[![gendiff demo](https://img.shields.io/badge/asciinema-gendiff-blue)](docs/gendiff-flat-json.cast)

Просмотреть локальную запись можно командой:

```bash
asciinema play docs/gendiff-flat-json.cast
```

---

<details>
<summary>Автоматические тесты Хекслета</summary>

Тесты запускаются на каждый коммит. За запуск отвечает файл `.github/workflows/hexlet-check.yml` — не удаляйте и не переименовывайте ни его, ни репозиторий.

</details>

## О Хекслете

[Хекслет](https://ru.hexlet.io/) — школа программирования: авторские программы обучения с практикой, поддержкой наставников и реальными проектами, которые остаются в резюме. Этот репозиторий — один из таких проектов.
