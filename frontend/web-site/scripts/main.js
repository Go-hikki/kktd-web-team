// Анимация загрузки
window.addEventListener('load', () => {
    const loader = document.querySelector('.loader');
    document.body.style.opacity = '1'; // Плавное появление контента
    setTimeout(() => {
        loader.classList.add('hidden');
        animateElements(); // Запускаем анимации после загрузки
    }, 500);
});

// Плавная прокрутка для якорных ссылок
document.querySelectorAll('a[href^="#"]').forEach(anchor => {
    anchor.addEventListener('click', (e) => {
        e.preventDefault();
        const targetId = anchor.getAttribute('href');
        if (targetId === '#') return;

        const targetElement = document.querySelector(targetId);
        if (targetElement) {
            window.scrollTo({
                top: targetElement.offsetTop - 60, // Учёт высоты хедера
                behavior: 'smooth'
            });
        }
        // Закрытие бургер-меню при клике на ссылку
        burgerMenu.classList.remove('active');
        navUl.classList.remove('active');
    });
});

// Эффект "липкого" хедера
const header = document.getElementById('main-header');
const scrollToTop = document.querySelector('.scroll-to-top');
window.addEventListener('scroll', () => {
    if (window.scrollY > 80) {
        header.classList.add('scrolled');
    } else {
        header.classList.remove('scrolled');
    }

    // Управление кнопкой "Наверх" только на мобильных (до 768px)
    if (window.innerWidth <= 768 && window.scrollY > 200) {
        scrollToTop.classList.add('visible');
    } else {
        scrollToTop.classList.remove('visible');
    }
});

// Прокрутка наверх при клике
scrollToTop.addEventListener('click', () => {
    window.scrollTo({ top: 0, behavior: 'smooth' });
});

// Анимация элементов при скролле
function animateElements() {
    const heroContent = document.querySelector('.hero-content');
    heroContent.classList.add('animated');

    const elements = document.querySelectorAll('.section-title, .program-card, .feature-item, .news-card, .footer-column');
    const observer = new IntersectionObserver(
        (entries) => {
            entries.forEach((entry) => {
                if (entry.isIntersecting) {
                    entry.target.classList.add('animated');
                }
            });
        },
        { threshold: 0.1 }
    );

    elements.forEach((element) => observer.observe(element));
}

// Модальное окно
const applyBtn = document.getElementById('apply-btn');
const applyModal = document.getElementById('apply-modal');
const closeModal = document.querySelector('.close-modal');

applyBtn.addEventListener('click', (e) => {
    e.preventDefault();
    applyModal.classList.add('active');
    document.body.style.overflow = 'hidden';
});

closeModal.addEventListener('click', () => {
    applyModal.classList.remove('active');
    document.body.style.overflow = '';
});

applyModal.addEventListener('click', (e) => {
    if (e.target === applyModal) {
        applyModal.classList.remove('active');
        document.body.style.overflow = '';
    }
});

// Обработка формы
const applicationForm = document.getElementById('application-form');
applicationForm.addEventListener('submit', (e) => {
    e.preventDefault();
    alert('Ваша заявка успешно отправлена! Мы свяжемся с вами в ближайшее время.');
    applyModal.classList.remove('active');
    document.body.style.overflow = '';
    applicationForm.reset();
});

// Анимация при наведении на кнопки
const buttons = document.querySelectorAll('.btn');
buttons.forEach((button) => {
    button.addEventListener('mouseenter', () => {
        button.style.transform = 'translateY(-3px)';
        button.style.boxShadow = '0 6px 12px rgba(0, 0, 0, 0.15)';
    });

    button.addEventListener('mouseleave', () => {
        button.style.transform = '';
        button.style.boxShadow = '';
    });
});

// Бургер-меню
const burgerMenu = document.querySelector('.burger-menu');
const navUl = document.querySelector('nav ul');

burgerMenu.addEventListener('click', () => {
    const isExpanded = burgerMenu.classList.toggle('active');
    navUl.classList.toggle('active');
    burgerMenu.setAttribute('aria-expanded', isExpanded);
});

// Закрытие бургер-меню при клике на ссылки
const navLinks = document.querySelectorAll('nav ul li a');
navLinks.forEach((link) => {
    link.addEventListener('click', () => {
        burgerMenu.classList.remove('active');
        navUl.classList.remove('active');
        burgerMenu.setAttribute('aria-expanded', 'false');
    });
});