// Анимация загрузки
window.addEventListener('load', function () {
    const loader = document.querySelector('.loader');
    setTimeout(function () {
        loader.classList.add('hidden');

        // Запускаем анимации после загрузки
        animateElements();
    }, 1000);
});

// Плавная прокрутка для якорных ссылок
document.querySelectorAll('a[href^="#"]').forEach(anchor => {
    anchor.addEventListener('click', function (e) {
        e.preventDefault();

        const targetId = this.getAttribute('href');
        if (targetId === '#') return;

        const targetElement = document.querySelector(targetId);
        if (targetElement) {
            window.scrollTo({
                top: targetElement.offsetTop - 60, // Уменьшено из-за компактного хедера
                behavior: 'smooth'
            });
        }
    });
});

// Эффект "липкого" хедера
const header = document.getElementById('main-header');
window.addEventListener('scroll', function () {
    if (window.scrollY > 80) { // Уменьшено
        header.classList.add('scrolled');
    } else {
        header.classList.remove('scrolled');
    }
});

// Анимация элементов при скролле
function animateElements() {
    const heroContent = document.querySelector('.hero-content');
    heroContent.classList.add('animated');

    const animateOnScroll = function () {
        const elements = document.querySelectorAll('.section-title, .program-card, .feature-item, .news-card, .footer-column');
        const windowHeight = window.innerHeight;

        elements.forEach(element => {
            const elementPosition = element.getBoundingClientRect().top;
            const animationPoint = windowHeight - 80;

            if (elementPosition < animationPoint) {
                element.classList.add('animated');
            }
        });
    };

    animateOnScroll();
    window.addEventListener('scroll', animateOnScroll);
}

// Модальное окно
const applyBtn = document.getElementById('apply-btn');
const applyModal = document.getElementById('apply-modal');
const closeModal = document.querySelector('.close-modal');

applyBtn.addEventListener('click', function (e) {
    e.preventDefault();
    applyModal.classList.add('active');
    document.body.style.overflow = 'hidden';
});

closeModal.addEventListener('click', function () {
    applyModal.classList.remove('active');
    document.body.style.overflow = '';
});

applyModal.addEventListener('click', function (e) {
    if (e.target === applyModal) {
        applyModal.classList.remove('active');
        document.body.style.overflow = '';
    }
});

// Обработка формы
const applicationForm = document.getElementById('application-form');
applicationForm.addEventListener('submit', function (e) {
    e.preventDefault();
    alert('Ваша заявка успешно отправлена! Мы свяжемся с вами в ближайшее время.');
    applyModal.classList.remove('active');
    document.body.style.overflow = '';
    applicationForm.reset();
});

// Анимация при наведении на кнопки
const buttons = document.querySelectorAll('.btn');
buttons.forEach(button => {
    button.addEventListener('mouseenter', function () {
        this.style.transform = 'translateY(-3px)';
        this.style.boxShadow = '0 6px 12px rgba(0, 0, 0, 0.15)';
    });

    button.addEventListener('mouseleave', function () {
        this.style.transform = '';
        this.style.boxShadow = '';
    });
});

// Плавное появление элементов при загрузке
setTimeout(() => {
    document.body.style.opacity = '1';
}, 100);

// Бургер-меню
const burgerMenu = document.querySelector('.burger-menu');
const navUl = document.querySelector('nav ul');

burgerMenu.addEventListener('click', () => {
    burgerMenu.classList.toggle('active');
    navUl.classList.toggle('active');
});

// Кнопка "Наверх"
const scrollToTopBtn = document.querySelector('.scroll-to-top');
scrollToTopBtn.addEventListener('click', () => {
    window.scrollTo({
        top: 0,
        behavior: 'smooth'
    });
});