

    1. Напишите функцию, которая на вход получает запрос SQL и произвольные параметры, среди которых могут быть как обычные значения (строка, число) так и слайсы таких значений.
    Позиция каждого переданного параметра в запросе SQL обозначается знаком "?".
    Функция должна вернуть запрос SQL, в котором для каждого параметра-слайса количество знаков "?" будет через запятую размножено до количества элементов слайса, а вторым ответом вернуть слайс из параметров, которые соответствуют новым позициям знаков "?".
    Пример:
    Вызов: func ( "SELECT * FROM table WHERE deleted = ? AND id IN(?) AND count < ?", false, []int{1, 6, 234}, 555 )
    Ответ: "SELECT * FROM table WHERE deleted = ? AND id IN(?,?,?) AND count < ?", []interface{}{ false, 1, 6, 234, 555 }

    2. Сделайте кодогенерацию с помощью easyjson для любой Вашей структуры.
