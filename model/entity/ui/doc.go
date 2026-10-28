// Package ui contains wrappers around Discord components (Components V1 and V2).
//
// Each component is a structure with JSON tags and a chain of builders. To ensure that the compiler catches incorrect embeddings (a button in a container, a text display in an ActionRow, etc.), components implement limiting interfaces.
// # Interfaces

// - Component — anything that can be placed in message.components.
// - UiV1 — something that can be placed in an ActionRow (V1 context).
// - Accessory — something that can be placed in Section.accessory.
// TopLvlComponent is a top—level component of message.components.
// Modal is something that can be put in a Label (modal windows).
//
// # Components and their interfaces
//
// ActionRow (1) Component, TopLvlComponent
// Button (2) Component, UiV1, Accessory
// SelectMenu (3,5,6,7,8) Component, UiV1, Modal (in Modal — only type 3)
// TextInput (4) Modal (only inside Label, only modals)
// Section (9) Component
// TextDisplay (10) Component
// Thumbnail (11) Component, Accessory
// MediaGallery (12) Component
// File (13) Component
// Separator (14) Component
// Container (17) Component, TopLvlComponent
// Label (18) Component, TopLvlComponent (Modal.components only)
//
// # Where can I put
//
// message.components : ActionRow, Container, Section, TextDisplay,
// Thumbnail, MediaGallery, File, Separator
// ActionRow.components : UiV1 — Button, SelectMenu
// Container.components : Component — ActionRow, Section, TextDisplay,
// Thumbnail, MediaGallery, File, Separator
// Section.components : only *textDisplay
// Section.accessory : Accessory — Button, Thumbnail
// Label.component : Modal — TextInput, SelectMenu (type 3 only)
// Modal.components : Label (only in MODAL type response)
//
// # Discord limitations that types do not catch
//
// - TextInput lives only in modal windows and it should be wrapped in Label,
// and not in ActionRow.
//   - Label cannot be placed in Container or ActionRow — only in Modal.
//   - Container requires the IS_COMPONENTS_V2 flag (1<<15) on the message.
//   - Container can contain ActionRow (V2 accepts V1, but not vice versa).
//   - In one ActionRow: up to 5 buttons OR one select — do not mix.
//   - SelectMenu in Label is allowed only with string options (type 3).
package ui
