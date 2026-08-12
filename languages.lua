-- Модуль тотальной локализации движка cppc+cgoluajshtml
local Localization = {}

-- Огромная база данных для всех языков мира
Localization.db = {
    ["RU"] = {
        create_project = "Создать новый проект",
        delete_project = "Удалить проект",
        open_code      = "Открыть код проекта",
        exit_editor    = "Выход из редактора",
        add_file       = "Добавить файл скрипта gd",
        spawn_3d       = "Заспавнить 3D блок",
        stretch_3d     = "Растянуть 3D блок по осям X, Y, Z",
        drag_png       = "Перетащить PNG и создать кнопку действия"
    },
    ["US"] = {
        create_project = "Create New Project",
        delete_project = "Delete Project",
        open_code      = "Open Project Code",
        exit_editor    = "Exit Editor",
        add_file       = "Add gd script file",
        spawn_3d       = "Spawn 3D block",
        stretch_3d     = "Stretch 3D block on X, Y, Z axes",
        drag_png       = "Drag PNG and create action button"
    },
    ["ZH"] = { -- Китайский язык
        create_project = "创建新项目",
        delete_project = "删除项目",
        open_code      = "打开项目 В源码",
        exit_editor    = "退出编辑器",
        add_file       = "添加 gd 脚本文件",
        spawn_3d       = "生成 3D 方块",
        stretch_3d     = "沿 X, Y, Z 轴拉伸 3D 方块",
        drag_png       = "拖放 PNG 并创建动作按钮"
    }
    -- Сюда можно дописать испанский, французский и абсолютно любой другой язык!
}

-- Текущий выбранный язык по умолчанию
local current_lang = "RU"

-- Функция смены языка (поддерживает любой код из базы)
function Localization.set_language(lang_code)
    if Localization.db[lang_code] then
        current_lang = lang_code
        print("[Lua Language System] Язык успешно переключен на: " .. lang_code)
    else
        print("[Lua Error] Этот язык пока не добавлен в базу, используем: " .. current_lang)
    end
end

-- Функция перевода конкретной кнопки
function Localization.translate(key)
    local lang_pack = Localization.db[current_lang]
    if lang_pack and lang_pack[key] then
        return lang_pack[key]
    else
        return "MISSING_TRANSLATION: " .. key
    end
end

-- Симуляция работы кнопок на разных языках
function Localization.test_all_languages()
    print("\n--- ТЕСТИРОВАНИЕ ВСЕХ ЯЗЫКОВ В ИНТЕРФЕЙСЕ (Lua) ---")
    
    -- Тест на Русском
    Localization.set_language("RU")
    print("Кнопка 1: " .. Localization.translate("create_project"))
    print("Кнопка 2: " .. Localization.translate("spawn_3d"))
    
    -- Тест на Английском
    Localization.set_language("US")
    print("Кнопка 1: " .. Localization.translate("create_project"))
    print("Кнопка 2: " .. Localization.translate("stretch_3d"))
    
    -- Тест на Китайском
    Localization.set_language("ZH")
    print("Кнопка 1: " .. Localization.translate("create_project"))
    print("Кнопка 3: " .. Localization.translate("drag_png"))
    print("---------------------------------------------------\n")
end

-- Запуск теста системы языков
Localization.test_all_languages()

return Localization
