export const MIN_SCALE = 0.25;
export const MAX_SCALE = 4;
export const ZOOM_STEP = 1.2;
export const MAX_URL_LENGTH = 8000;

export function clampScale(scale) {
    return Math.min(MAX_SCALE, Math.max(MIN_SCALE, scale));
}

export function isUrlTooLong(url) {
    return url.length > MAX_URL_LENGTH;
}

export class DiagramViewer {
    constructor(container) {
        this.container = container;
        this.image = null;
        this.baseWidth = 0;
        this.scale = 1;
        this.dragging = false;
        this.lastX = 0;
        this.lastY = 0;

        this.container.classList.add("diagram-viewer");

        this.container.addEventListener("wheel", (e) => {
            if (e.ctrlKey) {
                e.preventDefault();
                this.zoomAt(e.deltaY < 0 ? ZOOM_STEP : 1 / ZOOM_STEP, e.clientX, e.clientY);
            } else if (e.shiftKey) {
                e.preventDefault();
                this.container.scrollLeft += e.deltaX || e.deltaY;
            }
        }, { passive: false });

        this.container.addEventListener("mousedown", (e) => {
            this.dragging = true;
            this.lastX = e.clientX;
            this.lastY = e.clientY;
            this.container.classList.add("dragging");
        });

        document.addEventListener("mousemove", (e) => {
            if (!this.dragging) return;
            this.container.scrollLeft -= e.clientX - this.lastX;
            this.container.scrollTop -= e.clientY - this.lastY;
            this.lastX = e.clientX;
            this.lastY = e.clientY;
        });

        document.addEventListener("mouseup", () => {
            this.dragging = false;
            this.container.classList.remove("dragging");
        });
    }

    setImage(url, altText) {
        const image = document.createElement("img");
        image.alt = altText;
        image.draggable = false;
        this.image = image;

        image.addEventListener("load", () => {
            if (this.image !== image) return;
            this.baseWidth = image.naturalWidth;
            this.reset();
        });

        image.addEventListener("error", () => {
            if (this.image !== image) return;
            this.image = null;
            const error = document.createElement("div");
            error.className = "diagram-error";
            error.textContent = "Le diagramme n'a pas pu être chargé. Vérifiez votre connexion Internet.";
            this.container.replaceChildren(error);
        });

        if (isUrlTooLong(url)) {
            const warning = document.createElement("div");
            warning.className = "diagram-warning";
            warning.textContent = "L'algorithme est trop long pour être affiché : le diagramme pourrait ne pas se charger.";
            this.container.replaceChildren(warning, image);
        } else {
            this.container.replaceChildren(image);
        }

        image.src = url;
    }

    zoomIn() {
        this.zoomAt(ZOOM_STEP);
    }

    zoomOut() {
        this.zoomAt(1 / ZOOM_STEP);
    }

    reset() {
        this.scale = 1;
        this.applyScale();
        this.container.scrollLeft = 0;
        this.container.scrollTop = 0;
    }

    zoomAt(factor, clientX, clientY) {
        const image = this.image;
        if (!image || this.baseWidth === 0) return;

        const rect = this.container.getBoundingClientRect();
        const anchorX = clientX === undefined ? this.container.clientWidth / 2 : clientX - rect.left;
        const anchorY = clientY === undefined ? this.container.clientHeight / 2 : clientY - rect.top;

        const ratioX = (this.container.scrollLeft + anchorX - image.offsetLeft) / image.offsetWidth;
        const ratioY = (this.container.scrollTop + anchorY - image.offsetTop) / image.offsetHeight;

        this.scale = clampScale(this.scale * factor);
        this.applyScale();

        this.container.scrollLeft = image.offsetLeft + ratioX * image.offsetWidth - anchorX;
        this.container.scrollTop = image.offsetTop + ratioY * image.offsetHeight - anchorY;
    }

    applyScale() {
        if (!this.image || this.baseWidth === 0) return;
        this.image.style.width = `${this.baseWidth * this.scale}px`;
    }
}
