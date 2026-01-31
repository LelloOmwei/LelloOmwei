#[repr(C)]
pub struct OwPacket {
    pub magic: [u8; 2],    // Vždy "OW" (0x4F, 0x57)
    pub context: u16,      // ID ontológie (napr. 0x01 pre SmartHome)
    pub subject: u32,      // Ktoré zariadenie/entita to je
    pub predicate: u16,    // Čo sa deje (napr. 0x01 = meria)
    pub value: f32,        // Fyzikálna hodnota (v Kelvinoch!)
    pub timestamp: u32,    // Relatívny alebo Unix čas
}

impl OwPacket {
    pub fn new(context: u16, subject: u32, predicate: u16, value: f32) -> Self {
        Self {
            magic: *b"OW",
            context,
            subject,
            predicate,
            value,
            timestamp: 0, // Tu neskôr doplníme čas z hostiteľa
        }
    }
}
