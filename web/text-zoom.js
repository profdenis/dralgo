export const TEXT_MIN_SCALE = 0.5;
export const TEXT_MAX_SCALE = 3;
export const TEXT_ZOOM_STEP = 1.1;

export function clampTextScale(scale) {
    return Math.min(TEXT_MAX_SCALE, Math.max(TEXT_MIN_SCALE, scale));
}

export class TextZoom {
    constructor(element) {
        this.element = element;
        this.baseFontSize = parseFloat(getComputedStyle(element).fontSize);
        this.scale = 1;
        this.apply();

        this.element.addEventListener("wheel", (e) => {
            if (!e.ctrlKey) return;
            e.preventDefault();
            this.zoomBy(e.deltaY < 0 ? TEXT_ZOOM_STEP : 1 / TEXT_ZOOM_STEP);
        }, { passive: false });
    }

    zoomIn() {
        this.zoomBy(TEXT_ZOOM_STEP);
    }

    zoomOut() {
        this.zoomBy(1 / TEXT_ZOOM_STEP);
    }

    reset() {
        this.scale = 1;
        this.apply();
    }

    zoomBy(factor) {
        this.scale = clampTextScale(this.scale * factor);
        this.apply();
    }

    apply() {
        this.element.style.fontSize = `${this.baseFontSize * this.scale}px`;
    }
}
