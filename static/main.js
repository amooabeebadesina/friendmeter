Vue.component('name', {
    template: `<div class="name">
                <div class="name__title">
                    Please enter your name! (Example: Amy)
                </div>
                <div class="name__input">
                    <div class="control">
                        <input class="input is-rounded" type="text" placeholder="John">
                        <button class="button is-primary">Submit</button>
                    </div>
                </div>
            </div>`
});

window.Vue = new Vue({
    el: "#app"
})