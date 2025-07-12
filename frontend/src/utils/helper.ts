

export function formatDate(date: string): string {
    const spl = date.split('.')[0].split('T');
    return `${spl[0]} ${spl[1]}`
}