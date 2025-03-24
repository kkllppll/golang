document.addEventListener("DOMContentLoaded", function () {
    const form = document.querySelector("form");
    const container = document.querySelector(".container");

   

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
