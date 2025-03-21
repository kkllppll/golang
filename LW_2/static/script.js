document.addEventListener("DOMContentLoaded", function () {
    const form = document.querySelector("form");
    const fuelSelect = document.getElementById("fuelType");
    const selectedFuelText = document.getElementById("selectedFuel");
    const fuelTypeInput = document.getElementById("fuelTypeInput");

    // отримуємо вибраний тип палива з localStorage або встановлюємо за замовчуванням
    let selectedFuel = localStorage.getItem("selectedFuel") || "Вугілля";

    // вставляємо вибране паливо в інтерфейс
    selectedFuelText.textContent = selectedFuel;
    fuelTypeInput.value = selectedFuel;
    fuelSelect.value = selectedFuel;

    // оновлення localStorage при зміні вибору
    fuelSelect.addEventListener("change", function () {
        localStorage.setItem("selectedFuel", this.value);
        selectedFuelText.textContent = this.value;
        fuelTypeInput.value = this.value;
    });

    // обробка форми
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
                let resultContainer = document.querySelector(".result-container");
                resultContainer.innerHTML = newResult.innerHTML; // оновлення тільки результату
            }

            // прокручування до результатів після обчислення
            window.scrollTo({ top: 0, behavior: "smooth" });
        })
        .catch(error => console.error("Помилка запиту:", error));
    });
});
