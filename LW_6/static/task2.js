document.addEventListener("DOMContentLoaded", function () {
    const form = document.querySelector("form");
    const container = document.querySelector(".container");

    const addButton = document.createElement("button");
    addButton.textContent = "Додати елемент";
    addButton.type = "button";
    addButton.style.marginBottom = "10px";

    container.insertBefore(addButton, form);

    addButton.addEventListener("click", function () {
        addElementGroup();
    });

    function addElementGroup() {
        const group = document.createElement("div");
        group.classList.add("element-group");
        group.innerHTML = `
            <input type="text" name="epNameController" placeholder="Найменування ЕП" step="any" required>
            <input type="number" name="etaController[]" placeholder="Номінальний ККД (ηн)" step="any" required>
            <input type="number" name="cosPhiController[]" placeholder="Коефіцієнт потужності (cos φ)" step="any" required>
            <input type="number" name="voltageController[]" placeholder="Напруга (Uн)" step="any" required>
            <input type="number" name="quantityController[]" placeholder="Кількість ЕП" step="any" required>
            <input type="number" name="powerController[]" placeholder="Номінальна потужність (Pн)", step="any" required>
            <input type="number" name="usageCoeffController[]" placeholder="Коефіцієнт використання (Кв)" step="any" required>
            <input type="number" name="tgPhiController[]" placeholder="Коефіцієнт реактивної потужності (tg φ)" step="any" required>

            <button type="button" class="remove-btn">Видалити</button>
        `;

        form.insertBefore(group, form.querySelector("button[type='submit']"));

        group.querySelector(".remove-btn").addEventListener("click", function () {
            group.remove();
        });
    }

    form.addEventListener("submit", function (event) {
        event.preventDefault();

        const formData = new FormData(form);
        const params = new URLSearchParams();
        formData.forEach((value, key) => {
            params.append(key, value);
        });

        let actionUrl = window.location.pathname;

        fetch(actionUrl, {
            method: "POST",
            headers: {
                "Content-Type": "application/x-www-form-urlencoded"
            },
            body: params.toString()
        })
        .then(response => response.text())
        .then(html => {
            let parser = new DOMParser();
            let doc = parser.parseFromString(html, "text/html");
            let newResult = doc.querySelector(".result");

            if (newResult) {
                let oldResult = document.querySelector(".result");
                if (oldResult) {
                    oldResult.innerHTML = newResult.innerHTML; 
                } else {
                    container.appendChild(newResult);
                }
            }

            window.scrollTo({ top: 0, behavior: "smooth" });
        })
        .catch(error => console.error("Помилка запиту:", error));
    });
});
